import { RuleItem } from 'async-validator';
import { NodeTypeEnum } from './enumeration';
import { ISelectableParam } from '../../../workflow-expression-editor/model/data';
import {ICase, INode} from "@/api/dto/project";

type TriggerValue = 'blur' | 'change';

export interface CustomRuleItem extends RuleItem {
  trigger?: TriggerValue;
}

export type CustomRule = CustomRuleItem | CustomRuleItem[];

/**
 * 节点数据
 */
export interface IWorkflowNode {
  getName(): string;

  getType(): NodeTypeEnum;

  getIcon(): string;

  buildSelectableParam(nodeId: string): ISelectableParam | undefined;

  getFormRules(): Record<string, CustomRule>;

  /**
   * 校验
   * @throws Error
   */
  validate(): Promise<void>;
}

export interface IGlobal {
  concurrent: number | boolean;
}

/**
 * 工作流数据
 */
export interface IWorkflow {
  name: string;
  // description?: string;
  groupId: string;
  // global: IGlobal;
  // data: string;
  createTime: string;
  modifiedTime: string;
  cases?: ICase[];
  nodes?: INode[];
}

export type ValidateParamFn = (value: string) => void;

export type ValidateCacheFn = (name: string) => void;
