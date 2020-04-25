// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  Action,
  saveBuilder,
  saveBuilderX,
  WriteOperation,
  Changeset,
} from "ent/action";
import { Viewer, ID } from "ent/ent";
import Event from "src/ent/event";
import { EventBuilder, EventInput } from "src/ent/event/actions/event_builder";

export interface EventEditInput {
  name?: string;
  creatorID?: ID;
  startTime?: Date;
  endTime?: Date | null;
  location?: string;
}

export class EditEventActionBase implements Action<Event> {
  public readonly builder: EventBuilder;
  public readonly viewer: Viewer;
  private input: EventEditInput;

  constructor(viewer: Viewer, event: Event, input: EventEditInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new EventBuilder(
      this.viewer,
      WriteOperation.Edit,
      this,
      event,
    );
  }

  getInput(): EventInput {
    return {
      ...this.input,
      requiredFields: [],
    };
  }

  async changeset(): Promise<Changeset<Event>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<Event | null> {
    return await saveBuilder(this.builder);
  }

  async saveX(): Promise<Event> {
    return await saveBuilderX(this.builder);
  }

  static create<T extends EditEventActionBase>(
    this: new (viewer: Viewer, event: Event, input: EventEditInput) => T,
    viewer: Viewer,
    event: Event,
    input: EventEditInput,
  ): EditEventActionBase {
    return new this(viewer, event, input);
  }
}