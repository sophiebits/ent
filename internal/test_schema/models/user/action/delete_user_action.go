// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package action

import (
	"context"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models"
	builder "github.com/lolopinto/ent/internal/test_schema/models/user"
)

type DeleteUserAction struct {
	builder *builder.UserMutationBuilder
}

// DeleteUserFromContext is the factory method to get an ...
func DeleteUserFromContext(ctx context.Context, user *models.User) *DeleteUserAction {
	v, err := viewer.ForContext(ctx)
	if err != nil {
		panic("tried to perform mutation without a viewer")
	}
	return DeleteUser(v, user)
}

// DeleteUser is the factory method to get an ...
func DeleteUser(v viewer.ViewerContext, user *models.User) *DeleteUserAction {
	action := &DeleteUserAction{}
	builder := builder.NewMutationBuilder(
		v,
		ent.DeleteOperation,
		action.getFieldMap(),
		actions.ExistingEnt(user),
	)
	action.builder = builder
	return action
}

func (action *DeleteUserAction) GetBuilder() ent.MutationBuilder {
	return action.builder
}

func (action *DeleteUserAction) GetTypedBuilder() *builder.UserMutationBuilder {
	return action.builder
}

func (action *DeleteUserAction) GetViewer() viewer.ViewerContext {
	return action.builder.GetViewer()
}

func (action *DeleteUserAction) SetBuilderOnTriggers(triggers []actions.Trigger) error {
	return action.builder.SetTriggers(triggers)
}

func (action *DeleteUserAction) SetBuilderOnObservers(observers []actions.Observer) error {
	return action.builder.SetObservers(observers)
}

func (action *DeleteUserAction) GetChangeset() (ent.Changeset, error) {
	return actions.GetChangeset(action)
}

func (action *DeleteUserAction) Entity() ent.Entity {
	return nil
}

func (action *DeleteUserAction) ExistingEnt() ent.Entity {
	return action.builder.ExistingEnt()
}

// getFieldMap returns the fields that could be edited in this mutation
func (action *DeleteUserAction) getFieldMap() ent.ActionFieldMap {
	return ent.ActionFieldMap{}
}

// Validate returns an error if the current state of the action is not valid
func (action *DeleteUserAction) Validate() error {
	return action.builder.Validate()
}

// Save is the method called to execute this action and save change
func (action *DeleteUserAction) Save() error {
	return actions.Save(action)
}

var _ actions.Action = &DeleteUserAction{}
