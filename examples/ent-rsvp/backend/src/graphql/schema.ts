// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { GraphQLSchema } from "graphql";
import { QueryType } from "src/graphql/resolvers/generated/query_type";
import { MutationType } from "src/graphql/mutations/generated/mutation_type";
import {
  EventActivityRsvpStatusType,
  AddressType,
  EventActivityType,
  EventType,
  GuestGroupType,
  GuestType,
  UserType,
  EventActivityToAttendingConnectionType,
  EventActivityToDeclinedConnectionType,
  EventActivityToInvitesConnectionType,
  EventToEventActivitiesConnectionType,
  EventToGuestGroupsConnectionType,
  EventToGuestsConnectionType,
  GuestGroupToGuestsConnectionType,
  GuestGroupToInvitedEventsConnectionType,
  GuestToAttendingEventsConnectionType,
  GuestToDeclinedEventsConnectionType,
  UserToEventsConnectionType,
  ViewerTypeType,
} from "./resolvers";
import {
  ActivityEventCreateInput,
  EventCreateInputType,
  EventCreatePayloadType,
} from "src/graphql/mutations/generated/event/event_create_type";
import {
  AddressCreateInputType,
  AddressCreatePayloadType,
} from "src/graphql/mutations/generated/address/address_create_type";
import {
  AddressDeleteInputType,
  AddressDeletePayloadType,
} from "src/graphql/mutations/generated/address/address_delete_type";
import {
  AddressEditInputType,
  AddressEditPayloadType,
} from "src/graphql/mutations/generated/address/address_edit_type";
import {
  AddressEventActivityCreateInput,
  EventActivityCreateInputType,
  EventActivityCreatePayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_create_type";
import {
  EventActivityAddInviteInputType,
  EventActivityAddInvitePayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_add_invite_type";
import {
  EventActivityDeleteInputType,
  EventActivityDeletePayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_delete_type";
import {
  EventActivityEditInputType,
  EventActivityEditPayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_edit_type";
import {
  EventActivityRemoveInviteInputType,
  EventActivityRemoveInvitePayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_remove_invite_type";
import {
  EventActivityRsvpStatusEditInputType,
  EventActivityRsvpStatusEditPayloadType,
} from "src/graphql/mutations/generated/event_activity/event_activity_rsvp_status_edit_type";
import {
  GuestCreateInputType,
  GuestCreatePayloadType,
} from "src/graphql/mutations/generated/guest/guest_create_type";
import {
  GuestDeleteInputType,
  GuestDeletePayloadType,
} from "src/graphql/mutations/generated/guest/guest_delete_type";
import {
  GuestEditInputType,
  GuestEditPayloadType,
} from "src/graphql/mutations/generated/guest/guest_edit_type";
import {
  GuestGroupCreateInputType,
  GuestGroupCreatePayloadType,
  GuestGuestGroupCreateInput,
} from "src/graphql/mutations/generated/guest_group/guest_group_create_type";
import {
  GuestGroupDeleteInputType,
  GuestGroupDeletePayloadType,
} from "src/graphql/mutations/generated/guest_group/guest_group_delete_type";
import {
  GuestGroupEditInputType,
  GuestGroupEditPayloadType,
} from "src/graphql/mutations/generated/guest_group/guest_group_edit_type";
import {
  UserCreateInputType,
  UserCreatePayloadType,
} from "src/graphql/mutations/generated/user/user_create_type";

export default new GraphQLSchema({
  query: QueryType,
  mutation: MutationType,
  types: [
    EventActivityRsvpStatusType,
    AddressType,
    EventActivityType,
    EventType,
    GuestGroupType,
    GuestType,
    UserType,
    EventActivityToAttendingConnectionType(),
    EventActivityToDeclinedConnectionType(),
    EventActivityToInvitesConnectionType(),
    EventToEventActivitiesConnectionType(),
    EventToGuestGroupsConnectionType(),
    EventToGuestsConnectionType(),
    GuestGroupToGuestsConnectionType(),
    GuestGroupToInvitedEventsConnectionType(),
    GuestToAttendingEventsConnectionType(),
    GuestToDeclinedEventsConnectionType(),
    UserToEventsConnectionType(),
    ViewerTypeType,
    ActivityEventCreateInput,
    AddressCreateInputType,
    AddressCreatePayloadType,
    AddressDeleteInputType,
    AddressDeletePayloadType,
    AddressEditInputType,
    AddressEditPayloadType,
    AddressEventActivityCreateInput,
    EventActivityAddInviteInputType,
    EventActivityAddInvitePayloadType,
    EventActivityCreateInputType,
    EventActivityCreatePayloadType,
    EventActivityDeleteInputType,
    EventActivityDeletePayloadType,
    EventActivityEditInputType,
    EventActivityEditPayloadType,
    EventActivityRemoveInviteInputType,
    EventActivityRemoveInvitePayloadType,
    EventActivityRsvpStatusEditInputType,
    EventActivityRsvpStatusEditPayloadType,
    EventCreateInputType,
    EventCreatePayloadType,
    GuestCreateInputType,
    GuestCreatePayloadType,
    GuestDeleteInputType,
    GuestDeletePayloadType,
    GuestEditInputType,
    GuestEditPayloadType,
    GuestGroupCreateInputType,
    GuestGroupCreatePayloadType,
    GuestGroupDeleteInputType,
    GuestGroupDeletePayloadType,
    GuestGroupEditInputType,
    GuestGroupEditPayloadType,
    GuestGuestGroupCreateInput,
    UserCreateInputType,
    UserCreatePayloadType,
  ],
});