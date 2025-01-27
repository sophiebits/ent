// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  GraphQLFieldConfigMap,
  GraphQLID,
  GraphQLInt,
  GraphQLNonNull,
  GraphQLObjectType,
  GraphQLString,
} from "graphql";
import { RequestContext } from "@lolopinto/ent";
import {
  GraphQLEdgeConnection,
  GraphQLNodeInterface,
  nodeIDEncoder,
} from "@lolopinto/ent/graphql";
import { User, UserToEventsQuery } from "src/ent/";
import { UserToEventsConnectionType } from "src/graphql/resolvers/internal";

export const UserType = new GraphQLObjectType({
  name: "User",
  fields: (): GraphQLFieldConfigMap<User, RequestContext> => ({
    id: {
      type: GraphQLNonNull(GraphQLID),
      resolve: nodeIDEncoder,
    },
    firstName: {
      type: GraphQLNonNull(GraphQLString),
    },
    lastName: {
      type: GraphQLNonNull(GraphQLString),
    },
    emailAddress: {
      type: GraphQLNonNull(GraphQLString),
    },
    events: {
      type: GraphQLNonNull(UserToEventsConnectionType()),
      args: {
        first: {
          description: "",
          type: GraphQLInt,
        },
        after: {
          description: "",
          type: GraphQLString,
        },
        last: {
          description: "",
          type: GraphQLInt,
        },
        before: {
          description: "",
          type: GraphQLString,
        },
      },
      resolve: (user: User, args: {}, context: RequestContext) => {
        return new GraphQLEdgeConnection(
          user.viewer,
          user,
          (v, user: User) => UserToEventsQuery.query(v, user),
          args,
        );
      },
    },
  }),
  interfaces: [GraphQLNodeInterface],
  isTypeOf(obj) {
    return obj instanceof User;
  },
});
