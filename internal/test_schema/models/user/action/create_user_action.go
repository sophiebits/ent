// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package action

import (
	"context"
	"errors"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models"
	builder "github.com/lolopinto/ent/internal/test_schema/models/user"
)

type CreateUserAction struct {
	builder *builder.UserMutationBuilder
}

// CreateUserFromContext is the factory method to get an ...
func CreateUserFromContext(ctx context.Context) *CreateUserAction {
	v, err := viewer.ForContext(ctx)
	if err != nil {
		panic("tried to perform mutation without a viewer")
	}
	return CreateUser(v)
}

// CreateUser is the factory method to get an ...
func CreateUser(viewer viewer.ViewerContext) *CreateUserAction {
	action := &CreateUserAction{}
	builder := builder.NewMutationBuilder(
		viewer,
		ent.InsertOperation,
		action.getFieldMap(),
	)
	action.builder = builder
	return action
}

func (action *CreateUserAction) GetBuilder() ent.MutationBuilder {
	return action.builder
}

func (action *CreateUserAction) GetTypedBuilder() *builder.UserMutationBuilder {
	return action.builder
}

func (action *CreateUserAction) GetViewer() viewer.ViewerContext {
	return action.builder.GetViewer()
}

func (action *CreateUserAction) SetBuilderOnTriggers(triggers []actions.Trigger) error {
	action.builder.SetTriggers(triggers)
	for _, t := range triggers {
		trigger, ok := t.(builder.UserTrigger)
		if !ok {
			return errors.New("invalid trigger")
		}
		trigger.SetBuilder(action.builder)
	}
	return nil
}

func (action *CreateUserAction) GetChangeset() (ent.Changeset, error) {
	return actions.GetChangeset(action)
}

func (action *CreateUserAction) Entity() ent.Entity {
	return action.builder.GetUser()
}

func (action *CreateUserAction) ExistingEnt() ent.Entity {
	return action.builder.ExistingEnt()
}

// SetEmailAddress sets the EmailAddress while editing the User ent
func (action *CreateUserAction) SetEmailAddress(emailAddress string) *CreateUserAction {
	action.builder.SetEmailAddress(emailAddress)
	return action
}

// SetFirstName sets the FirstName while editing the User ent
func (action *CreateUserAction) SetFirstName(firstName string) *CreateUserAction {
	action.builder.SetFirstName(firstName)
	return action
}

// SetLastName sets the LastName while editing the User ent
func (action *CreateUserAction) SetLastName(lastName string) *CreateUserAction {
	action.builder.SetLastName(lastName)
	return action
}

// getFieldMap returns the fields that could be edited in this mutation
func (action *CreateUserAction) getFieldMap() ent.ActionFieldMap {
	return ent.ActionFieldMap{
		"EmailAddress": &ent.MutatingFieldInfo{
			DB:       "email_address",
			Required: true,
		},
		"FirstName": &ent.MutatingFieldInfo{
			DB:       "first_name",
			Required: true,
		},
		"LastName": &ent.MutatingFieldInfo{
			DB:       "last_name",
			Required: true,
		},
	}
}

// Validate returns an error if the current state of the action is not valid
func (action *CreateUserAction) Validate() error {
	return action.builder.Validate()
}

// Save is the method called to execute this action and save change
func (action *CreateUserAction) Save() (*models.User, error) {
	err := actions.Save(action)
	return action.builder.GetUser(), err
}

var _ actions.Action = &CreateUserAction{}
