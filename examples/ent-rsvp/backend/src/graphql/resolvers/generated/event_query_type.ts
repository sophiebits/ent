// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  GraphQLFieldConfig,
  GraphQLNonNull,
  GraphQLResolveInfo,
  GraphQLString,
} from "graphql";
import { RequestContext } from "@lolopinto/ent";
import { EventType } from "src/graphql/resolvers/internal";
import { EventResolver } from "../event";

export const EventQueryType: GraphQLFieldConfig<undefined, RequestContext> = {
  type: EventType,
  args: {
    slug: {
      description: "",
      type: GraphQLNonNull(GraphQLString),
    },
  },
  resolve: async (
    _source,
    { slug },
    context: RequestContext,
    _info: GraphQLResolveInfo,
  ) => {
    const r = new EventResolver();
    return r.event(context, slug);
  },
};
