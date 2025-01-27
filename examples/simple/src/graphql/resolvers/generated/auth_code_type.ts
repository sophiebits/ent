// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  GraphQLObjectType,
  GraphQLID,
  GraphQLString,
  GraphQLNonNull,
  GraphQLFieldConfig,
  GraphQLFieldConfigMap,
  GraphQLResolveInfo,
} from "graphql";
import { ID, RequestContext } from "@lolopinto/ent";
import { UserType } from "src/graphql/resolvers/";
import { AuthCode } from "src/ent/";

interface AuthCodeQueryArgs {
  id: ID;
}

export const AuthCodeType = new GraphQLObjectType({
  name: "AuthCode",
  fields: (): GraphQLFieldConfigMap<AuthCode, RequestContext> => ({
    user: {
      type: UserType,
      resolve: (authCode: AuthCode, args: {}) => {
        return authCode.loadUser();
      },
    },
    id: {
      type: GraphQLNonNull(GraphQLID),
    },
    code: {
      type: GraphQLNonNull(GraphQLString),
    },
    emailAddress: {
      type: GraphQLString,
    },
    phoneNumber: {
      type: GraphQLString,
    },
  }),
});

export const AuthCodeQuery: GraphQLFieldConfig<
  undefined,
  RequestContext,
  AuthCodeQueryArgs
> = {
  type: AuthCodeType,
  args: {
    id: {
      description: "",
      type: GraphQLNonNull(GraphQLID),
    },
  },
  resolve: async (
    _source,
    args: AuthCodeQueryArgs,
    context: RequestContext,
    _info: GraphQLResolveInfo,
  ) => {
    return AuthCode.load(context.getViewer(), args.id);
  },
};
