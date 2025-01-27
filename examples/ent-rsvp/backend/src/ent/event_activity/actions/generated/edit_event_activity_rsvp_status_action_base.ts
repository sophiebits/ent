// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerHasIdentityPrivacyPolicy,
  ID,
  PrivacyPolicy,
  Viewer,
} from "@lolopinto/ent";
import {
  Action,
  Changeset,
  WriteOperation,
  setEdgeTypeInGroup,
} from "@lolopinto/ent/action";
import { NodeType } from "src/ent";
import { EventActivity } from "src/ent/";
import {
  EventActivityBuilder,
  EventActivityInput,
} from "src/ent/event_activity/actions/event_activity_builder";

export enum EventActivityRsvpStatusInput {
  Attending = "attending",
  Declined = "declined",
}

export function getEventActivityRsvpStatusInputValues() {
  return [
    EventActivityRsvpStatusInput.Attending,
    EventActivityRsvpStatusInput.Declined,
  ];
}

export interface EditEventActivityRsvpStatusInput {
  rsvpStatus: EventActivityRsvpStatusInput;
  guestID: ID;
  dietaryRestrictions?: string | null;
}

export class EditEventActivityRsvpStatusActionBase
  implements Action<EventActivity>
{
  public readonly builder: EventActivityBuilder;
  public readonly viewer: Viewer;
  protected input: EditEventActivityRsvpStatusInput;
  protected eventActivity: EventActivity;

  constructor(
    viewer: Viewer,
    eventActivity: EventActivity,
    input: EditEventActivityRsvpStatusInput,
  ) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new EventActivityBuilder(
      this.viewer,
      WriteOperation.Edit,
      this,
      eventActivity,
    );
    this.eventActivity = eventActivity;
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): EventActivityInput {
    // we use a type assertion to override the weak type detection here
    return this.input as EventActivityInput;
  }

  async changeset(): Promise<Changeset<EventActivity>> {
    return this.builder.build();
  }

  private async setEdgeType() {
    await setEdgeTypeInGroup(
      this.builder.orchestrator,
      this.input.rsvpStatus,
      this.eventActivity.id,
      this.input.guestID,
      NodeType.EventActivity,
      this.eventActivity.getEventActivityRsvpStatusMap(),
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

  async save(): Promise<EventActivity | null> {
    await this.setEdgeType();
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<EventActivity> {
    await this.setEdgeType();
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends EditEventActivityRsvpStatusActionBase>(
    this: new (
      viewer: Viewer,
      eventActivity: EventActivity,
      input: EditEventActivityRsvpStatusInput,
    ) => T,
    viewer: Viewer,
    eventActivity: EventActivity,
    input: EditEventActivityRsvpStatusInput,
  ): EditEventActivityRsvpStatusActionBase {
    return new this(viewer, eventActivity, input);
  }

  static async saveXFromID<T extends EditEventActivityRsvpStatusActionBase>(
    this: new (
      viewer: Viewer,
      eventActivity: EventActivity,
      input: EditEventActivityRsvpStatusInput,
    ) => T,
    viewer: Viewer,
    id: ID,
    input: EditEventActivityRsvpStatusInput,
  ): Promise<EventActivity> {
    let eventActivity = await EventActivity.loadX(viewer, id);
    return await new this(viewer, eventActivity, input).saveX();
  }
}
