// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  GraphQLObjectType,
  GraphQLInputObjectType,
  GraphQLID,
  GraphQLString,
  GraphQLNonNull,
  GraphQLFieldConfig,
  GraphQLFieldConfigMap,
  GraphQLResolveInfo,
  GraphQLInputFieldConfigMap,
} from "graphql";
import { ID, RequestContext } from "@lolopinto/ent";
import { User } from "src/ent/";
import { UserType } from "src/graphql/resolvers/";
import EditEmailAddressAction, {
  EditEmailAddressInput,
} from "src/ent/user/actions/edit_email_address_action";

interface customEmailAddressEditInput extends EditEmailAddressInput {
  userID: ID;
}

interface EmailAddressEditResponse {
  user: User;
}

export const EmailAddressEditInputType = new GraphQLInputObjectType({
  name: "EmailAddressEditInput",
  fields: (): GraphQLInputFieldConfigMap => ({
    userID: {
      type: GraphQLNonNull(GraphQLID),
    },
    newEmail: {
      type: GraphQLNonNull(GraphQLString),
    },
  }),
});

export const EmailAddressEditResponseType = new GraphQLObjectType({
  name: "EmailAddressEditResponse",
  fields: (): GraphQLFieldConfigMap<
    EmailAddressEditResponse,
    RequestContext
  > => ({
    user: {
      type: GraphQLNonNull(UserType),
    },
  }),
});

export const EmailAddressEditType: GraphQLFieldConfig<
  undefined,
  RequestContext,
  { [input: string]: customEmailAddressEditInput }
> = {
  type: GraphQLNonNull(EmailAddressEditResponseType),
  args: {
    input: {
      description: "",
      type: GraphQLNonNull(EmailAddressEditInputType),
    },
  },
  resolve: async (
    _source,
    { input },
    context: RequestContext,
    _info: GraphQLResolveInfo,
  ): Promise<EmailAddressEditResponse> => {
    let user = await EditEmailAddressAction.saveXFromID(
      context.getViewer(),
      input.userID,
      {
        newEmail: input.newEmail,
      },
    );
    return { user: user };
  },
};
