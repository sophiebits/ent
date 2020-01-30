// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package action

import (
	"context"
	"time"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models"
	builder "github.com/lolopinto/ent/internal/test_schema/models/event"
)

type CreateEventAction struct {
	builder *builder.EventMutationBuilder
}

// CreateEventFromContext is the factory method to get an ...
func CreateEventFromContext(ctx context.Context) *CreateEventAction {
	v, err := viewer.ForContext(ctx)
	if err != nil {
		panic("tried to perform mutation without a viewer")
	}
	return CreateEvent(v)
}

// CreateEvent is the factory method to get an ...
func CreateEvent(v viewer.ViewerContext) *CreateEventAction {
	action := &CreateEventAction{}
	builder := builder.NewMutationBuilder(
		v,
		ent.InsertOperation,
		action.requiredFields(),
	)
	action.builder = builder
	return action
}

func (action *CreateEventAction) GetBuilder() ent.MutationBuilder {
	return action.builder
}

func (action *CreateEventAction) GetTypedBuilder() *builder.EventMutationBuilder {
	return action.builder
}

func (action *CreateEventAction) GetViewer() viewer.ViewerContext {
	return action.builder.GetViewer()
}

func (action *CreateEventAction) SetBuilderOnTriggers(triggers []actions.Trigger) {
	action.builder.SetTriggers(triggers)
}

func (action *CreateEventAction) SetBuilderOnObservers(observers []actions.Observer) {
	action.builder.SetObservers(observers)
}

func (action *CreateEventAction) SetBuilderOnValidators(validators []actions.Validator) {
	action.builder.SetValidators(validators)
}

func (action *CreateEventAction) GetChangeset() (ent.Changeset, error) {
	return actions.GetChangeset(action)
}

func (action *CreateEventAction) Entity() ent.Entity {
	return action.builder.GetEvent()
}

func (action *CreateEventAction) ExistingEnt() ent.Entity {
	return action.builder.ExistingEnt()
}

// SetName sets the Name while editing the Event ent
func (action *CreateEventAction) SetName(name string) *CreateEventAction {
	action.builder.SetName(name)
	return action
}

// SetUserID sets the UserID while editing the Event ent
func (action *CreateEventAction) SetUserID(userID string) *CreateEventAction {
	action.builder.SetUserID(userID)
	return action
}

// SetUserIDBuilder sets the UserID while editing the Event ent
func (action *CreateEventAction) SetUserIDBuilder(builder ent.MutationBuilder) *CreateEventAction {
	action.builder.SetUserIDBuilder(builder)
	return action
}

// SetStartTime sets the StartTime while editing the Event ent
func (action *CreateEventAction) SetStartTime(startTime time.Time) *CreateEventAction {
	action.builder.SetStartTime(startTime)
	return action
}

// SetEndTime sets the EndTime while editing the Event ent
func (action *CreateEventAction) SetEndTime(endTime time.Time) *CreateEventAction {
	action.builder.SetEndTime(endTime)
	return action
}

// SetNilableEndTime sets the EndTime while editing the Event ent
func (action *CreateEventAction) SetNilableEndTime(endTime *time.Time) *CreateEventAction {
	action.builder.SetNilableEndTime(endTime)
	return action
}

// SetLocation sets the Location while editing the Event ent
func (action *CreateEventAction) SetLocation(location string) *CreateEventAction {
	action.builder.SetLocation(location)
	return action
}

func (action *CreateEventAction) requiredFields() []string {
	return []string{
		"Name",
		"UserID",
		"StartTime",
		"Location",
	}
}

// Validate returns an error if the current state of the action is not valid
func (action *CreateEventAction) Validate() error {
	return action.builder.Validate()
}

// Save is the method called to execute this action and save change
func (action *CreateEventAction) Save() (*models.Event, error) {
	err := actions.Save(action)
	return action.builder.GetEvent(), err
}

var _ actions.Action = &CreateEventAction{}
