// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { GraphQLObjectType } from "graphql";
import {
  AddressQuery,
  AuthCodeQuery,
  ContactQuery,
  EventQuery,
  UserQuery,
  ViewerType,
} from "src/graphql/resolvers/";

export const QueryType = new GraphQLObjectType({
  name: "Query",
  fields: () => ({
    address: AddressQuery,
    authCode: AuthCodeQuery,
    contact: ContactQuery,
    event: EventQuery,
    user: UserQuery,
    viewer: ViewerType,
  }),
});
