// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerHasIdentityPrivacyPolicy,
  ID,
  PrivacyPolicy,
  Viewer,
} from "@lolopinto/ent";
import { Action, Changeset, WriteOperation } from "@lolopinto/ent/action";
import { User } from "src/ent/";
import { UserBuilder, UserInput } from "src/ent/user/actions/user_builder";

export interface EditEmailAddressInput {
  newEmail: string;
}

export class EditEmailAddressActionBase implements Action<User> {
  public readonly builder: UserBuilder;
  public readonly viewer: Viewer;
  protected input: EditEmailAddressInput;
  protected user: User;

  constructor(viewer: Viewer, user: User, input: EditEmailAddressInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new UserBuilder(
      this.viewer,
      WriteOperation.Edit,
      this,
      user,
    );
    this.user = user;
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): UserInput {
    // we use a type assertion to override the weak type detection here
    return this.input as UserInput;
  }

  async changeset(): Promise<Changeset<User>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<User | null> {
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<User> {
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends EditEmailAddressActionBase>(
    this: new (viewer: Viewer, user: User, input: EditEmailAddressInput) => T,
    viewer: Viewer,
    user: User,
    input: EditEmailAddressInput,
  ): EditEmailAddressActionBase {
    return new this(viewer, user, input);
  }

  static async saveXFromID<T extends EditEmailAddressActionBase>(
    this: new (viewer: Viewer, user: User, input: EditEmailAddressInput) => T,
    viewer: Viewer,
    id: ID,
    input: EditEmailAddressInput,
  ): Promise<User> {
    let user = await User.loadX(viewer, id);
    return await new this(viewer, user, input).saveX();
  }
}
