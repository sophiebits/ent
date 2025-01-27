// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  GraphQLFieldConfig,
  GraphQLFieldConfigMap,
  GraphQLID,
  GraphQLInputFieldConfigMap,
  GraphQLInputObjectType,
  GraphQLNonNull,
  GraphQLObjectType,
  GraphQLResolveInfo,
  GraphQLString,
} from "graphql";
import { RequestContext } from "@lolopinto/ent";
import { mustDecodeIDFromGQLID } from "@lolopinto/ent/graphql";
import { Address } from "src/ent/";
import EditAddressAction, {
  AddressEditInput,
} from "src/ent/address/actions/edit_address_action";
import { AddressType } from "src/graphql/resolvers/";

interface customAddressEditInput extends AddressEditInput {
  addressID: string;
  ownerID: string;
}

interface AddressEditPayload {
  address: Address;
}

export const AddressEditInputType = new GraphQLInputObjectType({
  name: "AddressEditInput",
  fields: (): GraphQLInputFieldConfigMap => ({
    addressID: {
      type: GraphQLNonNull(GraphQLID),
    },
    street: {
      type: GraphQLString,
    },
    city: {
      type: GraphQLString,
    },
    state: {
      type: GraphQLString,
    },
    zipCode: {
      type: GraphQLString,
    },
    apartment: {
      type: GraphQLString,
    },
    ownerID: {
      type: GraphQLID,
    },
    ownerType: {
      type: GraphQLString,
    },
  }),
});

export const AddressEditPayloadType = new GraphQLObjectType({
  name: "AddressEditPayload",
  fields: (): GraphQLFieldConfigMap<AddressEditPayload, RequestContext> => ({
    address: {
      type: GraphQLNonNull(AddressType),
    },
  }),
});

export const AddressEditType: GraphQLFieldConfig<
  undefined,
  RequestContext,
  { [input: string]: customAddressEditInput }
> = {
  type: GraphQLNonNull(AddressEditPayloadType),
  args: {
    input: {
      description: "",
      type: GraphQLNonNull(AddressEditInputType),
    },
  },
  resolve: async (
    _source,
    { input },
    context: RequestContext,
    _info: GraphQLResolveInfo,
  ): Promise<AddressEditPayload> => {
    let address = await EditAddressAction.saveXFromID(
      context.getViewer(),
      mustDecodeIDFromGQLID(input.addressID),
      {
        street: input.street,
        city: input.city,
        state: input.state,
        zipCode: input.zipCode,
        apartment: input.apartment,
      },
    );
    return { address: address };
  },
};
