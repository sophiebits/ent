export enum NodeType {
  // Address is the node type for the Address object. Used to identify this node in edges and other places.
  Address = "address",
  // Contact is the node type for the Contact object. Used to identify this node in edges and other places.
  Contact = "contact",
  // Event is the node type for the Event object. Used to identify this node in edges and other places.
  Event = "event",
  // User is the node type for the User object. Used to identify this node in edges and other places.
  User = "user",
}

export enum EdgeType {
  // EventToAttending is the edgeType for the event to attending edge.
  EventToAttending = "6ebc0c47-ea29-4635-b991-95e44162174d",
  // EventToDeclined is the edgeType for the event to declined edge.
  EventToDeclined = "db8d2454-f7b2-4147-aae1-e666daf3f3c3",
  // EventToHosts is the edgeType for the event to hosts edge.
  EventToHosts = "ebe3e709-845c-4723-ac9c-29f983f2b8ea",
  // EventToInvited is the edgeType for the event to invited edge.
  EventToInvited = "a72f5f64-3580-44fd-9bd0-d1335b803a46",
  // EventToMaybe is the edgeType for the event to maybe edge.
  EventToMaybe = "b0f6311b-fdab-4c26-b6bf-b751e0997735",
  // UserToCreatedEvents is the edgeType for the user to createdevents edge.
  UserToCreatedEvents = "daa3b2a3-8245-40ca-ae77-25bfb82578a7",
  // UserToDeclinedEvents is the edgeType for the user to declinedevents edge.
  UserToDeclinedEvents = "1c7c173b-63ce-4002-b121-4a87f82047dd",
  // UserToEventsAttending is the edgeType for the user to eventsattending edge.
  UserToEventsAttending = "2a98ba02-e342-4bb4-93f6-5d7ed02f5c48",
  // UserToFriends is the edgeType for the user to friends edge.
  UserToFriends = "d1a9316d-090f-4b02-b393-fd9372e2c905",
  // UserToInvitedEvents is the edgeType for the user to invitedevents edge.
  UserToInvitedEvents = "e439f2b2-d93a-4d1a-83f0-865bda5c8337",
  // UserToMaybeEvents is the edgeType for the user to maybeevents edge.
  UserToMaybeEvents = "8d5b1dee-ce65-452e-9f8d-78eca1993800",
}
