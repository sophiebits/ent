// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import { GraphQLSchema } from "graphql";
import {
  AddressCreateInputType,
  AddressCreatePayloadType,
} from "src/graphql/mutations/generated/address/address_create_type";
import {
  ContactCreateInputType,
  ContactCreatePayloadType,
} from "src/graphql/mutations/generated/contact/contact_create_type";
import {
  ContactDeleteInputType,
  ContactDeletePayloadType,
} from "src/graphql/mutations/generated/contact/contact_delete_type";
import {
  ContactEditInputType,
  ContactEditPayloadType,
} from "src/graphql/mutations/generated/contact/contact_edit_type";
import {
  EventAddHostInputType,
  EventAddHostPayloadType,
} from "src/graphql/mutations/generated/event/event_add_host_type";
import {
  EventCreateInputType,
  EventCreatePayloadType,
} from "src/graphql/mutations/generated/event/event_create_type";
import {
  EventDeleteInputType,
  EventDeletePayloadType,
} from "src/graphql/mutations/generated/event/event_delete_type";
import {
  EventEditInputType,
  EventEditPayloadType,
} from "src/graphql/mutations/generated/event/event_edit_type";
import {
  EventRemoveHostInputType,
  EventRemoveHostPayloadType,
} from "src/graphql/mutations/generated/event/event_remove_host_type";
import {
  EventRsvpStatusEditInputType,
  EventRsvpStatusEditPayloadType,
} from "src/graphql/mutations/generated/event/event_rsvp_status_edit_type";
import {
  HolidayCreateInputType,
  HolidayCreatePayloadType,
} from "src/graphql/mutations/generated/holiday/holiday_create_type";
import {
  HoursOfOperationCreateInputType,
  HoursOfOperationCreatePayloadType,
} from "src/graphql/mutations/generated/hours_of_operation/hours_of_operation_create_type";
import { MutationType } from "src/graphql/mutations/generated/mutation_type";
import {
  ConfirmEmailAddressEditInputType,
  ConfirmEmailAddressEditPayloadType,
} from "src/graphql/mutations/generated/user/confirm_email_address_edit_type";
import {
  ConfirmPhoneNumberEditInputType,
  ConfirmPhoneNumberEditPayloadType,
} from "src/graphql/mutations/generated/user/confirm_phone_number_edit_type";
import {
  EmailAddressEditInputType,
  EmailAddressEditPayloadType,
} from "src/graphql/mutations/generated/user/email_address_edit_type";
import {
  PhoneNumberEditInputType,
  PhoneNumberEditPayloadType,
} from "src/graphql/mutations/generated/user/phone_number_edit_type";
import {
  UserCreateInputType,
  UserCreatePayloadType,
} from "src/graphql/mutations/generated/user/user_create_type";
import {
  UserDeleteInputType,
  UserDeletePayloadType,
} from "src/graphql/mutations/generated/user/user_delete_type";
import {
  UserEditInputType,
  UserEditPayloadType,
} from "src/graphql/mutations/generated/user/user_edit_type";
import { QueryType } from "src/graphql/resolvers/generated/query_type";
import {
  AddressType,
  ContactType,
  EventRsvpStatusType,
  EventToAttendingConnectionType,
  EventToDeclinedConnectionType,
  EventToHostsConnectionType,
  EventToInvitedConnectionType,
  EventToMaybeConnectionType,
  EventType,
  GQLViewerType,
  HolidayType,
  HoursOfOperationType,
  UserToContactsConnectionType,
  UserToCreatedEventsConnectionType,
  UserToDeclinedEventsConnectionType,
  UserToEventsAttendingConnectionType,
  UserToFriendsConnectionType,
  UserToHostedEventsConnectionType,
  UserToInvitedEventsConnectionType,
  UserToMaybeEventsConnectionType,
  UserType,
  dayOfWeekType,
} from "./resolvers";

export default new GraphQLSchema({
  query: QueryType,
  mutation: MutationType,
  types: [
    EventRsvpStatusType,
    dayOfWeekType,
    AddressType,
    ContactType,
    EventType,
    HolidayType,
    HoursOfOperationType,
    UserType,
    EventToAttendingConnectionType(),
    EventToDeclinedConnectionType(),
    EventToHostsConnectionType(),
    EventToInvitedConnectionType(),
    EventToMaybeConnectionType(),
    UserToContactsConnectionType(),
    UserToCreatedEventsConnectionType(),
    UserToDeclinedEventsConnectionType(),
    UserToEventsAttendingConnectionType(),
    UserToFriendsConnectionType(),
    UserToHostedEventsConnectionType(),
    UserToInvitedEventsConnectionType(),
    UserToMaybeEventsConnectionType(),
    GQLViewerType,
    AddressCreateInputType,
    AddressCreatePayloadType,
    ConfirmEmailAddressEditInputType,
    ConfirmEmailAddressEditPayloadType,
    ConfirmPhoneNumberEditInputType,
    ConfirmPhoneNumberEditPayloadType,
    ContactCreateInputType,
    ContactCreatePayloadType,
    ContactDeleteInputType,
    ContactDeletePayloadType,
    ContactEditInputType,
    ContactEditPayloadType,
    EmailAddressEditInputType,
    EmailAddressEditPayloadType,
    EventAddHostInputType,
    EventAddHostPayloadType,
    EventCreateInputType,
    EventCreatePayloadType,
    EventDeleteInputType,
    EventDeletePayloadType,
    EventEditInputType,
    EventEditPayloadType,
    EventRemoveHostInputType,
    EventRemoveHostPayloadType,
    EventRsvpStatusEditInputType,
    EventRsvpStatusEditPayloadType,
    HolidayCreateInputType,
    HolidayCreatePayloadType,
    HoursOfOperationCreateInputType,
    HoursOfOperationCreatePayloadType,
    PhoneNumberEditInputType,
    PhoneNumberEditPayloadType,
    UserCreateInputType,
    UserCreatePayloadType,
    UserDeleteInputType,
    UserDeletePayloadType,
    UserEditInputType,
    UserEditPayloadType,
  ],
});
