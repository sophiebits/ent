// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
  Viewer,
} from "@lolopinto/ent";
import { Action, Changeset, WriteOperation } from "@lolopinto/ent/action";
import { Holiday } from "src/ent/";
import {
  HolidayBuilder,
  HolidayInput,
} from "src/ent/holiday/actions/holiday_builder";

export interface HolidayCreateInput {
  label: string;
  date: Date;
}

export class CreateHolidayActionBase implements Action<Holiday> {
  public readonly builder: HolidayBuilder;
  public readonly viewer: Viewer;
  protected input: HolidayCreateInput;

  constructor(viewer: Viewer, input: HolidayCreateInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new HolidayBuilder(this.viewer, WriteOperation.Insert, this);
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): HolidayInput {
    return this.input;
  }

  async changeset(): Promise<Changeset<Holiday>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<Holiday | null> {
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<Holiday> {
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends CreateHolidayActionBase>(
    this: new (viewer: Viewer, input: HolidayCreateInput) => T,
    viewer: Viewer,
    input: HolidayCreateInput,
  ): CreateHolidayActionBase {
    return new this(viewer, input);
  }
}
