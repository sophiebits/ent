import {
  ActionOperation,
  BaseEntSchema,
  Field,
  StringType,
  UUIDType,
  Action,
} from "@lolopinto/ent";

export default class Event extends BaseEntSchema {
  fields: Field[] = [
    StringType({ name: "Name" }),
    // start nullable so as to not break existing objects
    StringType({ name: "Slug", nullable: true, unique: true }),
    UUIDType({
      name: "creatorID",
      foreignKey: { schema: "User", column: "ID" },
    }),
  ];

  actions: Action[] = [
    {
      operation: ActionOperation.Create,
      actionOnlyFields: [
        {
          name: "activities",
          list: true,
          nullable: true,
          type: "Object",
          actionName: "CreateEventActivityAction",
        },
      ],
    },
    {
      operation: ActionOperation.Delete,
      hideFromGraphQL: true,
    },
  ];
}
