// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { GraphQLObjectType } from "graphql";
import { GraphQLConnectionType } from "@lolopinto/ent/graphql";
import { EventActivityToDeclinedEdge } from "src/ent/";
import { GuestType } from "src/graphql/resolvers/internal";

var connType: GraphQLConnectionType<
  GraphQLObjectType,
  EventActivityToDeclinedEdge
>;

export const EventActivityToDeclinedConnectionType = () => {
  if (connType === undefined) {
    connType = new GraphQLConnectionType("EventActivityToDeclined", GuestType);
  }
  return connType;
};
