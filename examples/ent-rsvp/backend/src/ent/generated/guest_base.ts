// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerPrivacyPolicy,
  AssocEdge,
  Context,
  Data,
  ID,
  LoadEntOptions,
  ObjectLoaderFactory,
  PrivacyPolicy,
  Viewer,
  loadEnt,
  loadEntX,
  loadEnts,
} from "@lolopinto/ent";
import { Field, getFields } from "@lolopinto/ent/schema";
import {
  EdgeType,
  Event,
  GuestGroup,
  GuestToAttendingEventsQuery,
  GuestToAuthCodesQuery,
  GuestToDeclinedEventsQuery,
  GuestToGuestDataQuery,
  NodeType,
} from "src/ent/internal";
import schema from "src/schema/guest";

const tableName = "guests";
const fields = [
  "id",
  "created_at",
  "updated_at",
  "name",
  "event_id",
  "email_address",
  "guest_group_id",
  "title",
];

export class GuestBase {
  readonly nodeType = NodeType.Guest;
  readonly id: ID;
  readonly createdAt: Date;
  readonly updatedAt: Date;
  readonly name: string;
  readonly eventID: ID;
  readonly emailAddress: string | null;
  readonly guestGroupID: ID;
  readonly title: string | null;

  constructor(public viewer: Viewer, data: Data) {
    this.id = data.id;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
    this.name = data.name;
    this.eventID = data.event_id;
    this.emailAddress = data.email_address;
    this.guestGroupID = data.guest_group_id;
    this.title = data.title;
  }

  // default privacyPolicy is Viewer can see themselves
  privacyPolicy: PrivacyPolicy = AllowIfViewerPrivacyPolicy;

  static async load<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    id: ID,
  ): Promise<T | null> {
    return loadEnt(viewer, id, GuestBase.loaderOptions.apply(this));
  }

  static async loadX<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    id: ID,
  ): Promise<T> {
    return loadEntX(viewer, id, GuestBase.loaderOptions.apply(this));
  }

  static async loadMany<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    ...ids: ID[]
  ): Promise<T[]> {
    return loadEnts(viewer, GuestBase.loaderOptions.apply(this), ...ids);
  }

  static async loadRawData<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
    id: ID,
    context?: Context,
  ): Promise<Data | null> {
    return await guestLoader.createLoader(context).load(id);
  }

  static async loadRawDataX<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
    id: ID,
    context?: Context,
  ): Promise<Data> {
    const row = await guestLoader.createLoader(context).load(id);
    if (!row) {
      throw new Error(`couldn't load row for ${id}`);
    }
    return row;
  }

  static loaderOptions<T extends GuestBase>(
    this: new (viewer: Viewer, data: Data) => T,
  ): LoadEntOptions<T> {
    return {
      tableName: tableName,
      fields: fields,
      ent: this,
      loaderFactory: guestLoader,
    };
  }

  private static schemaFields: Map<string, Field>;

  private static getSchemaFields(): Map<string, Field> {
    if (GuestBase.schemaFields != null) {
      return GuestBase.schemaFields;
    }
    return (GuestBase.schemaFields = getFields(schema));
  }

  static getField(key: string): Field | undefined {
    return GuestBase.getSchemaFields().get(key);
  }

  queryGuestToAttendingEvents(): GuestToAttendingEventsQuery {
    return GuestToAttendingEventsQuery.query(this.viewer, this.id);
  }

  queryGuestToDeclinedEvents(): GuestToDeclinedEventsQuery {
    return GuestToDeclinedEventsQuery.query(this.viewer, this.id);
  }

  queryAuthCodes(): GuestToAuthCodesQuery {
    return GuestToAuthCodesQuery.query(this.viewer, this.id);
  }

  queryGuestData(): GuestToGuestDataQuery {
    return GuestToGuestDataQuery.query(this.viewer, this.id);
  }

  async loadEvent(): Promise<Event | null> {
    return loadEnt(this.viewer, this.eventID, Event.loaderOptions());
  }

  loadEventX(): Promise<Event> {
    return loadEntX(this.viewer, this.eventID, Event.loaderOptions());
  }

  async loadGuestGroup(): Promise<GuestGroup | null> {
    return loadEnt(this.viewer, this.guestGroupID, GuestGroup.loaderOptions());
  }

  loadGuestGroupX(): Promise<GuestGroup> {
    return loadEntX(this.viewer, this.guestGroupID, GuestGroup.loaderOptions());
  }
}

export const guestLoader = new ObjectLoaderFactory({
  tableName,
  fields,
  key: "id",
});
