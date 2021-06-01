// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  Action,
  WriteOperation,
  Changeset,
  setEdgeTypeInGroup,
} from "@lolopinto/ent/action";
import {
  Viewer,
  ID,
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
} from "@lolopinto/ent";
import { NodeType } from "src/ent";
import { Event } from "src/ent/";
import { EventBuilder, EventInput } from "src/ent/event/actions/event_builder";

export enum EventRsvpStatusInput {
  Attending = "attending",
  Declined = "declined",
  Maybe = "maybe",
}

export function getEventRsvpStatusInputValues() {
  return [
    EventRsvpStatusInput.Attending,
    EventRsvpStatusInput.Declined,
    EventRsvpStatusInput.Maybe,
  ];
}

export interface EditEventRsvpStatusInput {
  rsvpStatus: EventRsvpStatusInput;
  userID: ID;
}

export class EditEventRsvpStatusActionBase implements Action<Event> {
  public readonly builder: EventBuilder;
  public readonly viewer: Viewer;
  protected input: EditEventRsvpStatusInput;
  protected event: Event;

  constructor(viewer: Viewer, event: Event, input: EditEventRsvpStatusInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new EventBuilder(
      this.viewer,
      WriteOperation.Edit,
      this,
      event,
    );
    this.event = event;
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): EventInput {
    // we use a type assertion to override the weak type detection here
    return this.input as EventInput;
  }

  async changeset(): Promise<Changeset<Event>> {
    return this.builder.build();
  }

  private async setEdgeType() {
    await setEdgeTypeInGroup(
      this.builder.orchestrator,
      this.input.rsvpStatus,
      this.event.id,
      this.input.userID,
      NodeType.Event,
      this.event.getEventRsvpStatusMap(),
    );
  }

  async valid(): Promise<boolean> {
    await this.setEdgeType();
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.setEdgeType();
    await this.builder.validX();
  }

  async save(): Promise<Event | null> {
    await this.setEdgeType();
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<Event> {
    await this.setEdgeType();
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends EditEventRsvpStatusActionBase>(
    this: new (
      viewer: Viewer,
      event: Event,
      input: EditEventRsvpStatusInput,
    ) => T,
    viewer: Viewer,
    event: Event,
    input: EditEventRsvpStatusInput,
  ): EditEventRsvpStatusActionBase {
    return new this(viewer, event, input);
  }

  static async saveXFromID<T extends EditEventRsvpStatusActionBase>(
    this: new (
      viewer: Viewer,
      event: Event,
      input: EditEventRsvpStatusInput,
    ) => T,
    viewer: Viewer,
    id: ID,
    input: EditEventRsvpStatusInput,
  ): Promise<Event> {
    let event = await Event.loadX(viewer, id);
    return await new this(viewer, event, input).saveX();
  }
}
