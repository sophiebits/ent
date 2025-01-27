// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { GraphQLObjectType } from "graphql";
import { GraphQLConnectionType } from "@lolopinto/ent/graphql";
import { EventToAttendingEdge } from "src/ent/";
import { UserType } from "src/graphql/resolvers/internal";

var connType: GraphQLConnectionType<GraphQLObjectType, EventToAttendingEdge>;

export const EventToAttendingConnectionType = () => {
  if (connType === undefined) {
    connType = new GraphQLConnectionType("EventToAttending", UserType);
  }
  return connType;
};
