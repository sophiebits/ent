// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { Action, WriteOperation, Changeset } from "@lolopinto/ent/action";
import {
  Viewer,
  ID,
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
} from "@lolopinto/ent";
import { EventActivity, GuestGroup } from "src/ent/";
import {
  EventActivityBuilder,
  EventActivityInput,
} from "src/ent/event_activity/actions/event_activity_builder";

export class EventActivityRemoveInviteActionBase
  implements Action<EventActivity> {
  public readonly builder: EventActivityBuilder;
  public readonly viewer: Viewer;
  protected eventActivity: EventActivity;

  constructor(viewer: Viewer, eventActivity: EventActivity) {
    this.viewer = viewer;
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
    return {};
  }

  removeInvite(...ids: ID[]): this;
  removeInvite(...nodes: GuestGroup[]): this;
  removeInvite(...nodes: ID[] | GuestGroup[]): this {
    nodes.forEach((node) => this.builder.removeInvite(node));
    return this;
  }
  async changeset(): Promise<Changeset<EventActivity>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<EventActivity | null> {
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<EventActivity> {
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends EventActivityRemoveInviteActionBase>(
    this: new (viewer: Viewer, eventActivity: EventActivity) => T,
    viewer: Viewer,
    eventActivity: EventActivity,
  ): EventActivityRemoveInviteActionBase {
    return new this(viewer, eventActivity);
  }

  static async saveXFromID<T extends EventActivityRemoveInviteActionBase>(
    this: new (viewer: Viewer, eventActivity: EventActivity) => T,
    viewer: Viewer,
    id: ID,
    inviteID: ID,
  ): Promise<EventActivity> {
    let eventActivity = await EventActivity.loadX(viewer, id);
    return await new this(viewer, eventActivity).removeInvite(inviteID).saveX();
  }
}
