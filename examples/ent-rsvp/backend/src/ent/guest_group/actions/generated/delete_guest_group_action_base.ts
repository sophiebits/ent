// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { Action, WriteOperation, Changeset } from "@lolopinto/ent/action";
import {
  Viewer,
  ID,
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
} from "@lolopinto/ent";
import { GuestGroup } from "src/ent/";
import {
  GuestGroupBuilder,
  GuestGroupInput,
} from "src/ent/guest_group/actions/guest_group_builder";

export class DeleteGuestGroupActionBase implements Action<GuestGroup> {
  public readonly builder: GuestGroupBuilder;
  public readonly viewer: Viewer;

  constructor(viewer: Viewer, guestGroup: GuestGroup) {
    this.viewer = viewer;
    this.builder = new GuestGroupBuilder(
      this.viewer,
      WriteOperation.Delete,
      this,
      guestGroup,
    );
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): GuestGroupInput {
    return {};
  }

  async changeset(): Promise<Changeset<GuestGroup>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<void> {
    await this.builder.save();
  }

  async saveX(): Promise<void> {
    await this.builder.saveX();
  }

  static create<T extends DeleteGuestGroupActionBase>(
    this: new (viewer: Viewer, guestGroup: GuestGroup) => T,
    viewer: Viewer,
    guestGroup: GuestGroup,
  ): DeleteGuestGroupActionBase {
    return new this(viewer, guestGroup);
  }

  static async saveXFromID<T extends DeleteGuestGroupActionBase>(
    this: new (viewer: Viewer, guestGroup: GuestGroup) => T,
    viewer: Viewer,
    id: ID,
  ): Promise<void> {
    let guestGroup = await GuestGroup.loadX(viewer, id);
    return await new this(viewer, guestGroup).saveX();
  }
}