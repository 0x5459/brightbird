package job

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/imdario/mergo"
	"github.com/ipfs-force-community/brightbird/env"

	"github.com/ipfs-force-community/brightbird/models"
	"github.com/ipfs-force-community/brightbird/repo"
	"github.com/ipfs-force-community/brightbird/types"
	"github.com/ipfs-force-community/brightbird/web/backend/config"
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	errors2 "k8s.io/apimachinery/pkg/api/errors"
)

var taskLog = logging.Logger("task")

type TaskMgr struct {
	c            *cron.Cron
	jobRepo      repo.IJobRepo
	taskRepo     repo.ITaskRepo
	testFlowRepo repo.ITestFlowRepo
	testRunner   *TestRunnerDeployer
	imageBuilder *ImageBuilderMgr

	privateRegistry types.PrivateRegistry
	runnerConfig    string
	cfg             config.Config
}

func NewTaskMgr(cfg config.Config, c *cron.Cron, jobRepo repo.IJobRepo, taskRepo repo.ITaskRepo, testFlowRepo repo.ITestFlowRepo, testRunner *TestRunnerDeployer, imageBuilder *ImageBuilderMgr, runnerConfig string, privateReg types.PrivateRegistry) *TaskMgr {
	return &TaskMgr{
		cfg:             cfg,
		c:               c,
		jobRepo:         jobRepo,
		taskRepo:        taskRepo,
		testFlowRepo:    testFlowRepo,
		testRunner:      testRunner,
		imageBuilder:    imageBuilder,
		runnerConfig:    runnerConfig,
		privateRegistry: privateReg,
	}
}

func (taskMgr *TaskMgr) Start(ctx context.Context) error {
	tm := time.NewTicker(time.Minute)
	defer tm.Stop()

	for {
		taskLog.Infof("loop check task status and start new task")
		//scan tasks to process
		jobs, err := taskMgr.jobRepo.List(ctx) //todo 数据规模大了 可以考虑采用被动触发的方式 现在这种做法简单一些
		if err != nil {
			taskLog.Error("fetch job list fail %v", err)
			continue
		}
		for _, job := range jobs {
			//try to remove finish testrunner
			err := taskMgr.testRunner.RemoveFinishRunner(ctx)
			if err != nil {
				taskLog.Errorf("clean finish scriptRunner %v", err)
				continue
			}
			//check running task state
			runningTask, err := taskMgr.taskRepo.List(ctx, models.PageReq[repo.ListTaskParams]{
				PageNum:  1,
				PageSize: math.MaxInt64,
				Params: repo.ListTaskParams{
					JobID: job.ID,
					State: []models.State{models.Running},
				},
			})
			if err != nil {
				taskLog.Errorf("fetch running task list fail %v", err)
				continue
			}

			for _, task := range runningTask.List {
				isClean := false
				if len(task.PodName) == 0 {
					//很少发生
					markFailErr := taskMgr.taskRepo.MarkState(ctx, task.ID, models.Error, "pod name is empty")
					if markFailErr != nil {
						log.Errorf("cannot mark task as fail %v origin err %v", err, markFailErr)
					}
					isClean = true
				}
				_, err = taskMgr.testRunner.CheckTestRunner(ctx, task.PodName)
				if err != nil {
					if errors2.IsNotFound(err) {
						markFailErr := taskMgr.taskRepo.MarkState(ctx, task.ID, models.Error, "failed 3 times, delete task")
						if markFailErr != nil {
							log.Errorf("cannot mark task as fail %v origin err %v", err, markFailErr)
						}
						isClean = true
					} else {
						log.Errorf("task id(%s) name(%s) try runner exceed more than 3 times, mark error and remove", task.ID, task.Name)
						// mark pod as fail and remove this pod
						markFailErr := taskMgr.taskRepo.MarkState(ctx, task.ID, models.Error, "failed 3 times, delete task")
						if markFailErr != nil {
							log.Errorf("cannot mark task as fail %v origin err %v", err, markFailErr)
						}
						isClean = true
					}
				}

				if isClean {
					cleanK8sErr := taskMgr.testRunner.CleanTestResource(ctx, string(task.TestId)) //clean again to ensure release all resource
					if cleanK8sErr != nil {
						log.Errorf("cannot clean k8s resource %v %v", cleanK8sErr)
					}
				}
			}

			// start init task
			initTasks, err := taskMgr.taskRepo.List(ctx, models.PageReq[repo.ListTaskParams]{
				PageNum:  1,
				PageSize: math.MaxInt64,
				Params: repo.ListTaskParams{
					JobID: job.ID,
					State: []models.State{models.Init, models.Building},
				},
			})
			if err != nil {
				taskLog.Errorf("fetch task list fail %v", err)
				continue
			}

			for _, task := range initTasks.List {
				pod, err := taskMgr.Process(ctx, task)
				if err != nil {
					taskLog.Errorf("process task fail %v", err)
					innerErr := taskMgr.taskRepo.MarkState(ctx, task.ID, models.Error, err.Error())
					if innerErr != nil {
						taskLog.Errorf("append log error %v", innerErr)
					}
					continue
				}

				err = taskMgr.taskRepo.UpdatePodRunning(ctx, task.ID, pod.Name)
				if err != nil {
					taskLog.Errorf("run task list fail %v", err)
				}
			}
		}

		select {
		case <-ctx.Done():
			return nil
		case <-tm.C:
		}
	}
}

func (taskMgr *TaskMgr) StopOneTask(ctx context.Context, id primitive.ObjectID) error {
	task, err := taskMgr.taskRepo.Get(ctx, &repo.GetTaskReq{ID: id})
	if err != nil {
		return err
	}

	err = taskMgr.testRunner.CleanTestResource(ctx, string(task.TestId))
	if err != nil {
		return err
	}

	task.State = models.Error
	task.Logs = append(task.Logs, "stop manually")
	_, err = taskMgr.taskRepo.Save(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func (taskMgr *TaskMgr) Process(ctx context.Context, task *models.Task) (*corev1.Pod, error) {
	task, err := taskMgr.taskRepo.IncreaseRetry(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	job, err := taskMgr.jobRepo.Get(ctx, task.JobId)
	if err != nil {
		return nil, err
	}

	testflow, err := taskMgr.testFlowRepo.Get(ctx, &repo.GetTestFlowParams{ID: job.TestFlowId})
	if err != nil {
		taskLog.Errorf("get test flow failed %v", err)
		return nil, err
	}
	graph := &models.Graph{}
	err = yaml.Unmarshal([]byte(testflow.Graph), graph)
	if err != nil {
		return nil, err
	}

	if task.BeforeBuild() {
		//update task state to build completed
		err = taskMgr.taskRepo.MarkState(ctx, task.ID, models.Building)
		if err != nil {
			return nil, err
		}
		//todo 不在runner里面做的原因 1. 编译需要比较好的性能，而runner可能会调度到比较差的机器上 2. 网络问题，拉代码编译过程很慢， 代理太费流量， 这里使用同一份代码可以缓解一些

		//confirm version and build image.
		taskLog.Infof("start to build image for testflow %s job %s", testflow.Name, job.Name)
		commitMap, err := taskMgr.imageBuilder.BuildTestFlowEnv(ctx, graph.Pipeline, task.InheritVersions) //todo maybe move this code to previous step
		if err != nil {
			return nil, err
		}

		var pipelines []*types.ExecNode
		for _, node := range graph.Pipeline {
			pipelines = append(pipelines, node.Value)
		}
		//save testflow as task params
		err = taskMgr.taskRepo.UpdatePipeline(ctx, task.ID, pipelines)
		if err != nil {
			return nil, err
		}

		//save testflow as task params
		err = taskMgr.taskRepo.UpdateCommitMap(ctx, task.ID, commitMap)
		if err != nil {
			return nil, err
		}
	}

	//run test flow
	file, err := os.Open(taskMgr.runnerConfig)
	if err != nil {
		return nil, err
	}

	var defaultGlobal = make(env.GlobalParams)
	defaultGlobal["logLevel"] = "DEBUG"

	//append global config
	err = mergo.Merge(&defaultGlobal, env.GlobalParams(taskMgr.cfg.CustomProperties))
	if err != nil {
		return nil, err
	}

	//append testflow params
	for _, value := range testflow.GlobalProperties {
		defaultGlobal[value.Name] = value.Value
	}

	//append job global params
	for _, value := range job.GlobalProperties {
		defaultGlobal[value.Name] = value.Value
	}

	//yaml escape character
	globalParamsBytes, err := json.Marshal(defaultGlobal)
	if err != nil {
		return nil, err
	}

	globalParamsBytes, err = yaml.Marshal(string(globalParamsBytes))
	if err != nil {
		return nil, err
	}

	//--log-level=DEBUG, --namespace={{.NameSpace}},--config=/shared-dir/config-template.toml, --plugins=/shared-dir/plugins, --taskId={{.TaskID}}
	args := fmt.Sprintf(`"--plugins=/shared-dir/plugins", "--namespace=%s", "--retry=%d", "--dbName=%s", "--mongoUrl=%s", "--mysql=%s", "--registry=%s", "--taskId=%s", --globalParams, %s`,
		taskMgr.cfg.NameSpace,
		task.RetryTime,
		taskMgr.cfg.DBName,
		taskMgr.cfg.MongoURL,
		taskMgr.cfg.Mysql,
		taskMgr.privateRegistry,
		task.ID.Hex(),
		string(globalParamsBytes),
	)

	log.Infof("invoke testrunner args %s", args)

	return taskMgr.testRunner.ApplyRunner(ctx, file, map[string]string{
		"NameSpace": taskMgr.cfg.NameSpace,
		"Registry":  string(taskMgr.privateRegistry),
		"TestID":    string(task.TestId),
		"ReTry":     strconv.Itoa(task.RetryTime),
		"Args":      args,
	})
}
