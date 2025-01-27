// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
  Viewer,
} from "@lolopinto/ent";
import { Action, Changeset, WriteOperation } from "@lolopinto/ent/action";
import { HoursOfOperation, dayOfWeek } from "src/ent/";
import {
  HoursOfOperationBuilder,
  HoursOfOperationInput,
} from "src/ent/hours_of_operation/actions/hours_of_operation_builder";

export interface HoursOfOperationCreateInput {
  dayOfWeek: dayOfWeek;
  open: Date;
  close: Date;
}

export class CreateHoursOfOperationActionBase
  implements Action<HoursOfOperation>
{
  public readonly builder: HoursOfOperationBuilder;
  public readonly viewer: Viewer;
  protected input: HoursOfOperationCreateInput;

  constructor(viewer: Viewer, input: HoursOfOperationCreateInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new HoursOfOperationBuilder(
      this.viewer,
      WriteOperation.Insert,
      this,
    );
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): HoursOfOperationInput {
    return this.input;
  }

  async changeset(): Promise<Changeset<HoursOfOperation>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<HoursOfOperation | null> {
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<HoursOfOperation> {
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends CreateHoursOfOperationActionBase>(
    this: new (viewer: Viewer, input: HoursOfOperationCreateInput) => T,
    viewer: Viewer,
    input: HoursOfOperationCreateInput,
  ): CreateHoursOfOperationActionBase {
    return new this(viewer, input);
  }
}
