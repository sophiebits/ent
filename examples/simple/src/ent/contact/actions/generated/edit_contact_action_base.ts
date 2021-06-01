// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  Action,
  Builder,
  WriteOperation,
  Changeset,
} from "@lolopinto/ent/action";
import {
  Viewer,
  ID,
  AllowIfViewerHasIdentityPrivacyPolicy,
  PrivacyPolicy,
} from "@lolopinto/ent";
import { Contact, User } from "src/ent/";
import {
  ContactBuilder,
  ContactInput,
} from "src/ent/contact/actions/contact_builder";

export interface ContactEditInput {
  emailAddress?: string;
  firstName?: string;
  lastName?: string;
  userID?: ID | Builder<User>;
}

export class EditContactActionBase implements Action<Contact> {
  public readonly builder: ContactBuilder;
  public readonly viewer: Viewer;
  protected input: ContactEditInput;
  protected contact: Contact;

  constructor(viewer: Viewer, contact: Contact, input: ContactEditInput) {
    this.viewer = viewer;
    this.input = input;
    this.builder = new ContactBuilder(
      this.viewer,
      WriteOperation.Edit,
      this,
      contact,
    );
    this.contact = contact;
  }

  getPrivacyPolicy(): PrivacyPolicy {
    return AllowIfViewerHasIdentityPrivacyPolicy;
  }

  getInput(): ContactInput {
    return this.input;
  }

  async changeset(): Promise<Changeset<Contact>> {
    return this.builder.build();
  }

  async valid(): Promise<boolean> {
    return this.builder.valid();
  }

  async validX(): Promise<void> {
    await this.builder.validX();
  }

  async save(): Promise<Contact | null> {
    await this.builder.save();
    return await this.builder.editedEnt();
  }

  async saveX(): Promise<Contact> {
    await this.builder.saveX();
    return await this.builder.editedEntX();
  }

  static create<T extends EditContactActionBase>(
    this: new (viewer: Viewer, contact: Contact, input: ContactEditInput) => T,
    viewer: Viewer,
    contact: Contact,
    input: ContactEditInput,
  ): EditContactActionBase {
    return new this(viewer, contact, input);
  }

  static async saveXFromID<T extends EditContactActionBase>(
    this: new (viewer: Viewer, contact: Contact, input: ContactEditInput) => T,
    viewer: Viewer,
    id: ID,
    input: ContactEditInput,
  ): Promise<Contact> {
    let contact = await Contact.loadX(viewer, id);
    return await new this(viewer, contact, input).saveX();
  }
}
