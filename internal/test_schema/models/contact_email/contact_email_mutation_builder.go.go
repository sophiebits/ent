// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package contact_email

import (
	"errors"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models"
	"github.com/lolopinto/ent/internal/test_schema/models/configs"
)

type ContactEmailMutationBuilder struct {
	builder          *actions.EntMutationBuilder
	contactEmail     *models.ContactEmail
	emailAddress     *string
	label            *string
	contactID        *string
	contactIDBuilder ent.MutationBuilder
}

func NewMutationBuilder(
	v viewer.ViewerContext,
	operation ent.WriteOperation,
	fieldMap ent.ActionFieldMap,
	opts ...func(*actions.EntMutationBuilder),
) *ContactEmailMutationBuilder {
	var contactEmail models.ContactEmail
	b := actions.NewMutationBuilder(
		v,
		operation,
		&contactEmail,
		&configs.ContactEmailConfig{},
		opts...,
	)
	b.FieldMap = fieldMap
	return &ContactEmailMutationBuilder{
		builder:      b,
		contactEmail: &contactEmail,
	}
}

func (b *ContactEmailMutationBuilder) SetEmailAddress(emailAddress string) *ContactEmailMutationBuilder {
	b.emailAddress = &emailAddress
	b.builder.SetField("EmailAddress", emailAddress)
	return b
}

func (b *ContactEmailMutationBuilder) SetLabel(label string) *ContactEmailMutationBuilder {
	b.label = &label
	b.builder.SetField("Label", label)
	return b
}

func (b *ContactEmailMutationBuilder) SetContactID(contactID string) *ContactEmailMutationBuilder {
	b.contactID = &contactID
	b.builder.SetField("ContactID", contactID)
	return b
}

func (b *ContactEmailMutationBuilder) SetContactIDBuilder(builder ent.MutationBuilder) *ContactEmailMutationBuilder {
	b.contactIDBuilder = builder
	b.builder.SetField("ContactID", builder)
	return b
}

func (b *ContactEmailMutationBuilder) GetEmailAddress() string {
	if b.emailAddress == nil {
		return ""
	}
	return *b.emailAddress
}

func (b *ContactEmailMutationBuilder) GetLabel() string {
	if b.label == nil {
		return ""
	}
	return *b.label
}

func (b *ContactEmailMutationBuilder) GetContactID() string {
	if b.contactID == nil {
		return ""
	}

	if b.contactIDBuilder != nil {
		return b.contactIDBuilder.GetPlaceholderID()
	}
	return *b.contactID
}

func (b *ContactEmailMutationBuilder) GetContactIDBuilder() ent.MutationBuilder {
	return b.contactIDBuilder
}

func (b *ContactEmailMutationBuilder) Validate() error {
	return b.builder.Validate()
}

func (b *ContactEmailMutationBuilder) GetViewer() viewer.ViewerContext {
	return b.builder.GetViewer()
}

func (b *ContactEmailMutationBuilder) GetContactEmail() *models.ContactEmail {
	return b.contactEmail
}

func (b *ContactEmailMutationBuilder) SetTriggers(triggers []actions.Trigger) error {
	b.builder.SetTriggers(triggers)
	for _, t := range triggers {
		trigger, ok := t.(ContactEmailTrigger)
		if !ok {
			return errors.New("invalid trigger")
		}
		trigger.SetBuilder(b)
	}
	return nil
}

// SetObservers sets the builder on an observer. Unlike SetTriggers, it's not required that observers implement the ContactEmailObserver
// interface since there's expected to be more reusability here e.g. generic logging, generic send text observer etc
func (b *ContactEmailMutationBuilder) SetObservers(observers []actions.Observer) error {
	b.builder.SetObservers(observers)
	for _, o := range observers {
		observer, ok := o.(ContactEmailObserver)
		if ok {
			observer.SetBuilder(b)
		}
	}
	return nil
}

func (b *ContactEmailMutationBuilder) GetChangeset() (ent.Changeset, error) {
	return b.builder.GetChangeset()
}

func (b *ContactEmailMutationBuilder) ExistingEnt() ent.Entity {
	return b.builder.ExistingEnt()
}

func (b *ContactEmailMutationBuilder) Entity() ent.Entity {
	return b.builder.Entity()
}

func (b *ContactEmailMutationBuilder) GetOperation() ent.WriteOperation {
	return b.builder.GetOperation()
}

func (b *ContactEmailMutationBuilder) GetPlaceholderID() string {
	return b.builder.GetPlaceholderID()
}

var _ ent.MutationBuilder = &ContactEmailMutationBuilder{}

type ContactEmailTrigger interface {
	SetBuilder(*ContactEmailMutationBuilder)
}

type ContactEmailMutationBuilderTrigger struct {
	Builder *ContactEmailMutationBuilder
}

func (trigger *ContactEmailMutationBuilderTrigger) SetBuilder(b *ContactEmailMutationBuilder) {
	trigger.Builder = b
}

type ContactEmailObserver interface {
	SetBuilder(*ContactEmailMutationBuilder)
}

type ContactEmailMutationBuilderObserver struct {
	Builder *ContactEmailMutationBuilder
}

func (observer *ContactEmailMutationBuilderObserver) SetBuilder(b *ContactEmailMutationBuilder) {
	observer.Builder = b
}