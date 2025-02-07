<template>
  <div class="project-group" v-loading="loading">
    <folding :status="toggle">
      <template #prefix>
        <span class="prefix-wrapper">
          <i
            :class="['jm-icon-button-right', 'prefix', toggle ? 'rotate' : '']"
            :disabled="projectPage.total === 0"
            @click="saveFoldStatus(toggle, testflowGroup?.id)"
          />
        </span>
      </template>
      <template #title>
        <div class="title">
          <div class="left">
            <div class="group-name">
              {{ testflowGroup?.name }}
              <div class="description">
                {{ ('('+ (testflowGroup?.description || '无') + ')' )  }}
              </div>
            </div>
          </div>
          <div class="right">
            <span class="desc"
              >（共有
              {{ projectPage.total >= 0 ? projectPage.total : 0 }}
              个测试流）</span
            >
            <div class="update-time">
              <span>最后修改时间：</span
              ><span>{{ datetimeFormatter(testflowGroup?.modifiedTime) }}</span>
            </div>
            <div class="operation">
                <div
                  class="edit op-item"
                  @click="
                    toEdit(
                      testflowGroup?.id,
                      testflowGroup?.name,
                      testflowGroup?.isShow,
                      testflowGroup?.description
                    )
                  "
                ></div>
                <div
                  class="delete op-item"
                  @click="toDelete(testflowGroup?.name, testflowGroup?.id)"
                ></div>
              </div>
          </div>
        </div>
      </template>
      <template #default>
        <div>
          <div class="testflows" v-show="toggle && testflows.length > 0">
            <jm-empty
              description="暂无测试流"
              :image-size="98"
              v-if="testflows.length === 0"
            />
            <project-item
              v-else
              v-for="project of testflows"
              :key="project.id"
              :project="project"
              @triggered="handleProjectTriggered"
              @synchronized="handleProjectSynchronized"
              @deleted="handleProjectDeleted"
              @terminated="handleProjectTerminated"
            />
          </div>
        </div>
      </template>
    </folding>
  </div>
  <group-editor
        :name="groupName || ''"
        :description="groupDescription"
        :is-show="showInHomePage"
        :id="groupId || ''"
        v-if="editionActivated"
        @closed="editionActivated = false"
        @completed="editCompleted"
      />
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  getCurrentInstance,
  nextTick,
  onBeforeMount,
  onUpdated,
  PropType,
  ref,
} from 'vue';
import { ITestFlowDetail } from '@/api/dto/testflow';
import { ITestflowGroupVo } from '@/api/dto/testflow-group';
import { queryTestFlow } from '@/api/view-no-auth';
import { IQueryForm } from '@/model/modules/project';
import ProjectItem from '@/views/common/project-item.vue';
import { IPageVo } from '@/api/dto/common';
import { Mutable } from '@/utils/lib';
import { START_PAGE_NUM } from '@/utils/constants';
import Folding from '@/views/common/folding.vue';
import { createNamespacedHelpers, useStore } from 'vuex';
import { namespace } from '@/store/modules/project-group';
import noDataImg from '@/assets/svgs/index/no-data.svg';
import sleep from '@/utils/sleep';
import { datetimeFormatter } from '@/utils/formatter';
import GroupEditor from '@/views/project-group/project-group-editor.vue';
import { eventBus } from '@/main';
import { deleteProjectGroup } from '@/api/testflow-group';

export default defineComponent({
  components: { ProjectItem, Folding, GroupEditor },
  props: {
    // 测试流组
    testflowGroup: {
      type: Object as PropType<ITestflowGroupVo>,
    },
    // 查询关键字
    name: {
      type: String,
    },
    // 是否开启移动模式
    move: {
      type: Boolean,
      default: false,
    },
    projectName:{
      type:String,
      require:true,
    },
  },
  setup(props: any) {
    const store = useStore();
    const { mapMutations } = createNamespacedHelpers(namespace);
    const groupName = ref<string>();
    const editionActivated = ref<boolean>(false);
    const groupId = ref<string>();
    const groupDescription = ref<string>();
    const showInHomePage = ref<boolean>(false);
    const projectGroupFoldingMapping = store.state[namespace];
    // 根据测试流组在vuex中保存的状态，进行展开、折叠间的切换
    const toggle = computed<boolean>(() => {
      // 只有全等于为undefined说明该测试流组一开始根本没有做折叠操作
      if (projectGroupFoldingMapping[props.testflowGroup?.id] === undefined) {
        return true;
      }
      return projectGroupFoldingMapping[props.testflowGroup.id];
    });

    const { proxy } = getCurrentInstance() as any;
    const loading = ref<boolean>(false);
    const projectPage = ref<Mutable<IPageVo<ITestFlowDetail>>>({
      total: -1,
      pages: 0,
      list: [],
      pageNum: START_PAGE_NUM,
    });
    const testflows = computed<ITestFlowDetail[]>(() => projectPage.value.list.filter(value=>{
      if (props?.projectName !== '') {
        return value.name.includes(props?.projectName);
      }
      return true;
    }));

    const queryForm = ref<IQueryForm>({
      pageNum: START_PAGE_NUM,
      pageSize: 40,
      groupId: props.testflowGroup?.id,
      name: props.name,
    });
    // 保存单个测试流组的展开折叠状态
    const saveFoldStatus = (status: boolean, id?: string) => {
      // 改变状态
      const toggle = !status;
      // 调用vuex的mutations更改对应测试流组的状态
      proxy.mutate({
        id,
        status: toggle,
      });
    };
    // 重新加载当前已经加载过的测试流
    const reloadCurrentProjectList = async () => {
      try {
        const { pageSize, pageNum } = queryForm.value;
        // 获得当前已经加载了的总数
        const currentCount = pageSize * pageNum;
        projectPage.value = await queryTestFlow({
          ...queryForm.value,
          pageNum: START_PAGE_NUM,
          pageSize: currentCount,
        });
        console.log(projectPage.value);
      } catch (err) {
        proxy.$throw(err, proxy);
      }
    };

    const loadProject = async () => {
      projectPage.value = await queryTestFlow({ ...queryForm.value });
      // 测试流组中测试流为空，将其自动折叠
      if (projectPage.value.total === 0) {
        saveFoldStatus(true, props.testflowGroup?.id);
      }
      return;
    };
    const toEdit = (
      id?: string,
      name?: string,
      isShow?: boolean,
      description?: string,
    ) => {
      groupName.value = name;
      groupDescription.value = description;
      groupId.value = id;
      showInHomePage.value = isShow ?? false;
      editionActivated.value = true;      
    };
    const toDelete = async (name?: string, groupId?: string) => {
      let msg = '<div>确定要删除分组吗?</div>';
      msg += `<div style="margin-top: 5px; font-size: 12px; line-height: normal;">名称：${name}</div>`;

      proxy
        .$confirm(msg, '删除分组', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          dangerouslyUseHTMLString: true,
        })
        .then(async () => {
          if (!groupId) { return; }
          try {
            await deleteProjectGroup(groupId);            
            proxy.$success('测试流分组删除成功');
            eventBus.emit('newGroup');
          } catch (err) {
            proxy.$throw(err, proxy);
          }
        })
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        .catch(() => {
        });
    };
    const editCompleted = () => {
      eventBus.emit('newGroup');
    };
    // 初始化测试流列表
    onBeforeMount(async () => {
      await nextTick(() => {
        queryForm.value.name = props.name;
      });
      await loadProject();
    });
    onUpdated(async () => {
      if (
        queryForm.value.name === props.name &&
        queryForm.value.groupId === props.testflowGroup?.id
      ) {
        return;
      }
      queryForm.value.name = props.name;
      queryForm.value.groupId = props.testflowGroup?.id;
    });

    reloadCurrentProjectList(); // init
    return {
      editCompleted,
      toEdit,
      toDelete,
      datetimeFormatter,
      noDataImg,
      ...mapMutations({
        mutate: 'mutate',
      }),
      toggle,
      loading,
      projectPage,
      testflows,
      queryForm,
      handleProjectSynchronized: async () => {
        // 刷新测试流列表，保留查询状态
        await reloadCurrentProjectList();
      },
      handleProjectDeleted: (id: string) => {
        // const index = testflows.value.findIndex(item => item.id === id);
        loadProject();        
      },
      handleProjectTriggered: async (id: string) => {
        await sleep(400);
        // 刷新测试流列表，保留查询状态
        await reloadCurrentProjectList();
      },
      handleProjectTerminated: async (id: string) => {
        // 刷新测试流列表，保留查询状态
        await reloadCurrentProjectList();
      },
      saveFoldStatus,
      projectGroupFoldingMapping,
      groupName,
      editionActivated,
      groupId,
      groupDescription,
      showInHomePage,
    };
  },
});
</script>

<style scoped lang="less">
.project-group {
  margin-top: 24px;

  .prefix-wrapper {
    display: flex;
    align-items: center;
    cursor: not-allowed;
  }

  .prefix {
    color: #6b7b8d;
    font-size: 12px;
    cursor: pointer;
    transition: all 0.1s linear;

    &[disabled="true"] {
      color: #a7b0bb;
      pointer-events: none;
    }

    &.rotate {
      transition: all 0.1s linear;
      transform: rotate(90deg);
    }
  }

  .title {
    display: flex;
    align-items: center;
    justify-content: space-between;
    .left {
      display: flex;
      align-items: flex-end;
      justify-content: space-between;
      margin-left: 10px;
      padding-right: 0.7%;
      color: #082340;
      font-weight: bold;
      font-size: 18px;

      .group-name {
        display: flex;
        align-items: center;
        .description {
          padding-left: 12px;
          color: #082340;
          font-weight: normal;
          font-size: 14px;
          opacity: 0.46;
        }
      }
      .more-container {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 86px;
        height: 24px;
        border-radius: 15px;
        background: #eff7ff;
        font-weight: 400;
        font-size: 12px;
        cursor: pointer;

        a {
          color: #6b7b8d;
          line-height: 24px;
        }

        .more-icon {
          position: relative;
          top: 1.4px;
          right: 0px;
          display: inline-block;
          width: 12px;
          height: 12px;
          background: url("@/assets/svgs/btn/more.svg") no-repeat;
          text-align: center;
          line-height: 12px;
        }

        &:hover {
          color: #096dd9;

          a {
            color: #096dd9;
          }

          .more-icon {
            background: url("@/assets/svgs/btn/more-active.svg") no-repeat;
          }
        }
      }
    }

    .right {
      display: flex;
      margin-left: 12px;
      color: #082340;
      font-weight: normal;
      font-size: 14px;
      opacity: 0.46;

      .operation {
            display: flex;
            margin-left: 5px;

            .op-item {
              width: 22px;
              height: 22px;
              background-size: contain;
              cursor: pointer;

              &:active {
                border-radius: 4px;
                background-color: #eff7ff;
              }

              &.edit {
                background-image: url('@/assets/svgs/btn/edit.svg');
              }

              &.delete {
                margin-left: 15px;
                background-image: url('@/assets/svgs/btn/del.svg');
              }
            }
          }
    }
  }

  .testflows {
    display: flex;
    flex-wrap: wrap;

    .el-empty {
      padding-top: 102px;
    }

    ::v-deep(.jm-sorter) {
      .drag-target-insertion {
        // width: v-bind(spacing);
        border-width: 0;
        background-color: transparent;

        &::after {
          position: absolute;
          top: 0;
          left: 20%;
          box-sizing: border-box;
          width: 60%;
          height: 100%;
          border: 2px solid #096dd9;
          background: rgba(9, 109, 217, 0.3);
          content: "";
        }
      }
    }

    .list {
      display: flex;
      flex-wrap: wrap;
      width: 100%;
    }
  }

  .load-more {
    display: flex;
    justify-content: center;
    margin: 10px auto 0px;
  }
}
</style>
