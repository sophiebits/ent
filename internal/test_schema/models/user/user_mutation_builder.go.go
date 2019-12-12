// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package user

import (
	"errors"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models"
	"github.com/lolopinto/ent/internal/test_schema/models/configs"
)

type UserMutationBuilder struct {
	builder      *actions.EntMutationBuilder
	user         *models.User
	emailAddress *string
	firstName    *string
	lastName     *string
	bio          *string
}

func NewMutationBuilder(
	v viewer.ViewerContext,
	operation ent.WriteOperation,
	fieldMap ent.ActionFieldMap,
	opts ...func(*actions.EntMutationBuilder),
) *UserMutationBuilder {
	var user models.User
	b := actions.NewMutationBuilder(
		v,
		operation,
		&user,
		&configs.UserConfig{},
		opts...,
	)
	b.FieldMap = fieldMap
	return &UserMutationBuilder{
		builder: b,
		user:    &user,
	}
}

func (b *UserMutationBuilder) SetEmailAddress(emailAddress string) *UserMutationBuilder {
	b.emailAddress = &emailAddress
	b.builder.SetField("EmailAddress", emailAddress)
	return b
}

func (b *UserMutationBuilder) SetFirstName(firstName string) *UserMutationBuilder {
	b.firstName = &firstName
	b.builder.SetField("FirstName", firstName)
	return b
}

func (b *UserMutationBuilder) SetLastName(lastName string) *UserMutationBuilder {
	b.lastName = &lastName
	b.builder.SetField("LastName", lastName)
	return b
}

func (b *UserMutationBuilder) SetBio(bio string) *UserMutationBuilder {
	b.bio = &bio
	b.builder.SetField("Bio", bio)
	return b
}

func (b *UserMutationBuilder) SetNilableBio(bio *string) *UserMutationBuilder {
	b.bio = bio
	if bio == nil {
		b.builder.SetField("Bio", nil)
	} else {
		b.builder.SetField("Bio", *bio)
	}
	return b
}

func (b *UserMutationBuilder) GetEmailAddress() string {
	if b.emailAddress == nil {
		return ""
	}
	return *b.emailAddress
}

func (b *UserMutationBuilder) GetFirstName() string {
	if b.firstName == nil {
		return ""
	}
	return *b.firstName
}

func (b *UserMutationBuilder) GetLastName() string {
	if b.lastName == nil {
		return ""
	}
	return *b.lastName
}

func (b *UserMutationBuilder) GetBio() *string {
	if b.bio == nil {
		return nil
	}
	return b.bio
}

// AddEvents adds one or more instances of Event to the Events edge while editing the Event ent
func (b *UserMutationBuilder) AddEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.AddEventID(event.ID)
	}
	return b
}

// AddEventIDs adds an instance of Event to the Events edge while editing the Event ent
func (b *UserMutationBuilder) AddEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.AddEventID(eventID)
	}
	return b
}

// AddEventID adds an instance of Event to the Events edge while editing the Event ent
func (b *UserMutationBuilder) AddEventID(eventID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToEventsEdge, eventID, models.EventType, options...)
	return b
}

// AddFamilyMembers adds one or more instances of User to the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) AddFamilyMembers(users ...*models.User) *UserMutationBuilder {
	for _, user := range users {
		b.AddFamilyMemberID(user.ID)
	}
	return b
}

// AddFamilyMemberIDs adds an instance of User to the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) AddFamilyMemberIDs(userIDs ...string) *UserMutationBuilder {
	for _, userID := range userIDs {
		b.AddFamilyMemberID(userID)
	}
	return b
}

// AddFamilyMemberID adds an instance of User to the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) AddFamilyMemberID(userID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToFamilyMembersEdge, userID, models.UserType, options...)
	return b
}

// AddFriends adds one or more instances of User to the Friends edge while editing the User ent
func (b *UserMutationBuilder) AddFriends(users ...*models.User) *UserMutationBuilder {
	for _, user := range users {
		b.AddFriendID(user.ID)
	}
	return b
}

// AddFriendIDs adds an instance of User to the Friends edge while editing the User ent
func (b *UserMutationBuilder) AddFriendIDs(userIDs ...string) *UserMutationBuilder {
	for _, userID := range userIDs {
		b.AddFriendID(userID)
	}
	return b
}

// AddFriendID adds an instance of User to the Friends edge while editing the User ent
func (b *UserMutationBuilder) AddFriendID(userID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToFriendsEdge, userID, models.UserType, options...)
	return b
}

// AddInvitedEvents adds one or more instances of Event to the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddInvitedEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.AddInvitedEventID(event.ID)
	}
	return b
}

// AddInvitedEventIDs adds an instance of Event to the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddInvitedEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.AddInvitedEventID(eventID)
	}
	return b
}

// AddInvitedEventID adds an instance of Event to the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddInvitedEventID(eventID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToInvitedEventsEdge, eventID, models.EventType, options...)
	return b
}

// AddEventsAttending adds one or more instances of Event to the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) AddEventsAttending(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.AddEventsAttendingID(event.ID)
	}
	return b
}

// AddEventsAttendingIDs adds an instance of Event to the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) AddEventsAttendingIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.AddEventsAttendingID(eventID)
	}
	return b
}

// AddEventsAttendingID adds an instance of Event to the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) AddEventsAttendingID(eventID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToEventsAttendingEdge, eventID, models.EventType, options...)
	return b
}

// AddDeclinedEvents adds one or more instances of Event to the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddDeclinedEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.AddDeclinedEventID(event.ID)
	}
	return b
}

// AddDeclinedEventIDs adds an instance of Event to the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddDeclinedEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.AddDeclinedEventID(eventID)
	}
	return b
}

// AddDeclinedEventID adds an instance of Event to the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) AddDeclinedEventID(eventID string, options ...func(*ent.EdgeOperation)) *UserMutationBuilder {
	b.builder.AddOutboundEdge(models.UserToDeclinedEventsEdge, eventID, models.EventType, options...)
	return b
}

// RemoveEvents removes an instance of Event from the Events edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.RemoveEventID(event.ID)
	}
	return b
}

// RemoveEventIDs removes an instance of Event from the Events edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.RemoveEventID(eventID)
	}
	return b
}

// RemoveEventID removes an instance of Event from the Events edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEventID(eventID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToEventsEdge, eventID, models.EventType)
	return b
}

// RemoveFamilyMembers removes an instance of User from the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) RemoveFamilyMembers(users ...*models.User) *UserMutationBuilder {
	for _, user := range users {
		b.RemoveFamilyMemberID(user.ID)
	}
	return b
}

// RemoveFamilyMemberIDs removes an instance of User from the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) RemoveFamilyMemberIDs(userIDs ...string) *UserMutationBuilder {
	for _, userID := range userIDs {
		b.RemoveFamilyMemberID(userID)
	}
	return b
}

// RemoveFamilyMemberID removes an instance of User from the FamilyMembers edge while editing the User ent
func (b *UserMutationBuilder) RemoveFamilyMemberID(userID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToFamilyMembersEdge, userID, models.UserType)
	return b
}

// RemoveFriends removes an instance of User from the Friends edge while editing the User ent
func (b *UserMutationBuilder) RemoveFriends(users ...*models.User) *UserMutationBuilder {
	for _, user := range users {
		b.RemoveFriendID(user.ID)
	}
	return b
}

// RemoveFriendIDs removes an instance of User from the Friends edge while editing the User ent
func (b *UserMutationBuilder) RemoveFriendIDs(userIDs ...string) *UserMutationBuilder {
	for _, userID := range userIDs {
		b.RemoveFriendID(userID)
	}
	return b
}

// RemoveFriendID removes an instance of User from the Friends edge while editing the User ent
func (b *UserMutationBuilder) RemoveFriendID(userID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToFriendsEdge, userID, models.UserType)
	return b
}

// RemoveInvitedEvents removes an instance of Event from the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveInvitedEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.RemoveInvitedEventID(event.ID)
	}
	return b
}

// RemoveInvitedEventIDs removes an instance of Event from the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveInvitedEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.RemoveInvitedEventID(eventID)
	}
	return b
}

// RemoveInvitedEventID removes an instance of Event from the InvitedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveInvitedEventID(eventID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToInvitedEventsEdge, eventID, models.EventType)
	return b
}

// RemoveEventsAttending removes an instance of Event from the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEventsAttending(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.RemoveEventsAttendingID(event.ID)
	}
	return b
}

// RemoveEventsAttendingIDs removes an instance of Event from the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEventsAttendingIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.RemoveEventsAttendingID(eventID)
	}
	return b
}

// RemoveEventsAttendingID removes an instance of Event from the EventsAttending edge while editing the Event ent
func (b *UserMutationBuilder) RemoveEventsAttendingID(eventID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToEventsAttendingEdge, eventID, models.EventType)
	return b
}

// RemoveDeclinedEvents removes an instance of Event from the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveDeclinedEvents(events ...*models.Event) *UserMutationBuilder {
	for _, event := range events {
		b.RemoveDeclinedEventID(event.ID)
	}
	return b
}

// RemoveDeclinedEventIDs removes an instance of Event from the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveDeclinedEventIDs(eventIDs ...string) *UserMutationBuilder {
	for _, eventID := range eventIDs {
		b.RemoveDeclinedEventID(eventID)
	}
	return b
}

// RemoveDeclinedEventID removes an instance of Event from the DeclinedEvents edge while editing the Event ent
func (b *UserMutationBuilder) RemoveDeclinedEventID(eventID string) *UserMutationBuilder {
	b.builder.RemoveOutboundEdge(models.UserToDeclinedEventsEdge, eventID, models.EventType)
	return b
}

func (b *UserMutationBuilder) Validate() error {
	return b.builder.Validate()
}

func (b *UserMutationBuilder) GetViewer() viewer.ViewerContext {
	return b.builder.GetViewer()
}

func (b *UserMutationBuilder) GetUser() *models.User {
	return b.user
}

func (b *UserMutationBuilder) SetTriggers(triggers []actions.Trigger) error {
	b.builder.SetTriggers(triggers)
	for _, t := range triggers {
		trigger, ok := t.(UserTrigger)
		if !ok {
			return errors.New("invalid trigger")
		}
		trigger.SetBuilder(b)
	}
	return nil
}

// SetObservers sets the builder on an observer. Unlike SetTriggers, it's not required that observers implement the UserObserver
// interface since there's expected to be more reusability here e.g. generic logging, generic send text observer etc
func (b *UserMutationBuilder) SetObservers(observers []actions.Observer) error {
	b.builder.SetObservers(observers)
	for _, o := range observers {
		observer, ok := o.(UserObserver)
		if ok {
			observer.SetBuilder(b)
		}
	}
	return nil
}

func (b *UserMutationBuilder) GetChangeset() (ent.Changeset, error) {
	return b.builder.GetChangeset()
}

func (b *UserMutationBuilder) ExistingEnt() ent.Entity {
	return b.builder.ExistingEnt()
}

func (b *UserMutationBuilder) Entity() ent.Entity {
	return b.builder.Entity()
}

func (b *UserMutationBuilder) GetOperation() ent.WriteOperation {
	return b.builder.GetOperation()
}

func (b *UserMutationBuilder) GetPlaceholderID() string {
	return b.builder.GetPlaceholderID()
}

var _ ent.MutationBuilder = &UserMutationBuilder{}

type UserTrigger interface {
	SetBuilder(*UserMutationBuilder)
}

type UserMutationBuilderTrigger struct {
	Builder *UserMutationBuilder
}

func (trigger *UserMutationBuilderTrigger) SetBuilder(b *UserMutationBuilder) {
	trigger.Builder = b
}

type UserObserver interface {
	SetBuilder(*UserMutationBuilder)
}

type UserMutationBuilderObserver struct {
	Builder *UserMutationBuilder
}

func (observer *UserMutationBuilderObserver) SetBuilder(b *UserMutationBuilder) {
	observer.Builder = b
}