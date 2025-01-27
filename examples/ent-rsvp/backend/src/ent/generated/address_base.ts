// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  AllowIfViewerPrivacyPolicy,
  Context,
  Data,
  Ent,
  ID,
  LoadEntOptions,
  ObjectLoaderFactory,
  PrivacyPolicy,
  Viewer,
  loadEnt,
  loadEntViaKey,
  loadEntX,
  loadEntXViaKey,
  loadEnts,
} from "@lolopinto/ent";
import { Field, getFields } from "@lolopinto/ent/schema";
import { NodeType } from "src/ent/internal";
import { loadEntByType, loadEntXByType } from "src/ent/loadAny";
import schema from "src/schema/address";

const tableName = "addresses";
const fields = [
  "id",
  "created_at",
  "updated_at",
  "street",
  "city",
  "state",
  "zip_code",
  "apartment",
  "owner_id",
  "owner_type",
];

export class AddressBase {
  readonly nodeType = NodeType.Address;
  readonly id: ID;
  readonly createdAt: Date;
  readonly updatedAt: Date;
  readonly street: string;
  readonly city: string;
  readonly state: string;
  readonly zipCode: string;
  readonly apartment: string | null;
  readonly ownerID: ID;
  readonly ownerType: string;

  constructor(public viewer: Viewer, data: Data) {
    this.id = data.id;
    this.createdAt = data.created_at;
    this.updatedAt = data.updated_at;
    this.street = data.street;
    this.city = data.city;
    this.state = data.state;
    this.zipCode = data.zip_code;
    this.apartment = data.apartment;
    this.ownerID = data.owner_id;
    this.ownerType = data.owner_type;
  }

  // default privacyPolicy is Viewer can see themselves
  privacyPolicy: PrivacyPolicy = AllowIfViewerPrivacyPolicy;

  static async load<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    id: ID,
  ): Promise<T | null> {
    return loadEnt(viewer, id, AddressBase.loaderOptions.apply(this));
  }

  static async loadX<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    id: ID,
  ): Promise<T> {
    return loadEntX(viewer, id, AddressBase.loaderOptions.apply(this));
  }

  static async loadMany<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    ...ids: ID[]
  ): Promise<T[]> {
    return loadEnts(viewer, AddressBase.loaderOptions.apply(this), ...ids);
  }

  static async loadRawData<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    id: ID,
    context?: Context,
  ): Promise<Data | null> {
    return await addressLoader.createLoader(context).load(id);
  }

  static async loadRawDataX<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    id: ID,
    context?: Context,
  ): Promise<Data> {
    const row = await addressLoader.createLoader(context).load(id);
    if (!row) {
      throw new Error(`couldn't load row for ${id}`);
    }
    return row;
  }

  static async loadFromOwnerID<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    ownerID: ID,
  ): Promise<T | null> {
    return loadEntViaKey(viewer, ownerID, {
      ...AddressBase.loaderOptions.apply(this),
      loaderFactory: addressOwnerIDLoader,
    });
  }

  static async loadFromOwnerIDX<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    viewer: Viewer,
    ownerID: ID,
  ): Promise<T> {
    return loadEntXViaKey(viewer, ownerID, {
      ...AddressBase.loaderOptions.apply(this),
      loaderFactory: addressOwnerIDLoader,
    });
  }

  static async loadIDFromOwnerID<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    ownerID: ID,
    context?: Context,
  ): Promise<ID | undefined> {
    const row = await addressOwnerIDLoader.createLoader(context).load(ownerID);
    return row?.id;
  }

  static async loadRawDataFromOwnerID<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
    ownerID: ID,
    context?: Context,
  ): Promise<Data | null> {
    return await addressOwnerIDLoader.createLoader(context).load(ownerID);
  }

  static loaderOptions<T extends AddressBase>(
    this: new (viewer: Viewer, data: Data) => T,
  ): LoadEntOptions<T> {
    return {
      tableName: tableName,
      fields: fields,
      ent: this,
      loaderFactory: addressLoader,
    };
  }

  private static schemaFields: Map<string, Field>;

  private static getSchemaFields(): Map<string, Field> {
    if (AddressBase.schemaFields != null) {
      return AddressBase.schemaFields;
    }
    return (AddressBase.schemaFields = getFields(schema));
  }

  static getField(key: string): Field | undefined {
    return AddressBase.getSchemaFields().get(key);
  }

  async loadOwner(): Promise<Ent | null> {
    return loadEntByType(
      this.viewer,
      this.ownerType as unknown as NodeType,
      this.ownerID,
    );
  }

  loadOwnerX(): Promise<Ent> {
    return loadEntXByType(
      this.viewer,
      this.ownerType as unknown as NodeType,
      this.ownerID,
    );
  }
}

export const addressLoader = new ObjectLoaderFactory({
  tableName,
  fields,
  key: "id",
});

export const addressOwnerIDLoader = new ObjectLoaderFactory({
  tableName,
  fields,
  key: "owner_id",
});

addressLoader.addToPrime(addressOwnerIDLoader);
addressOwnerIDLoader.addToPrime(addressLoader);
