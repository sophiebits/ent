// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"context"

	"github.com/lolopinto/ent/ent/test_schema/models"
)

type Resolver struct{}

func (r *Resolver) Contact() ContactResolver {
	return &contactResolver{r}
}
func (r *Resolver) Event() EventResolver {
	return &eventResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type contactResolver struct{ *Resolver }

func (r *contactResolver) AllowList(ctx context.Context, obj *models.Contact) ([]*models.User, error) {
	return obj.LoadAllowList()
}

type eventResolver struct{ *Resolver }

func (r *eventResolver) Invited(ctx context.Context, obj *models.Event) ([]*models.User, error) {
	return obj.LoadInvited()
}

func (r *eventResolver) User(ctx context.Context, obj *models.Event) (*models.User, error) {
	return obj.LoadUser()
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Contact(ctx context.Context, id string) (*models.Contact, error) {
	return models.LoadContactFromContext(ctx, id)
}

func (r *queryResolver) Event(ctx context.Context, id string) (*models.Event, error) {
	return models.LoadEventFromContext(ctx, id)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return models.LoadUserFromContext(ctx, id)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Contacts(ctx context.Context, obj *models.User) ([]*models.Contact, error) {
	return obj.LoadContacts()
}

func (r *userResolver) Events(ctx context.Context, obj *models.User) ([]*models.Event, error) {
	return obj.LoadEvents()
}

func (r *userResolver) FamilyMembers(ctx context.Context, obj *models.User) ([]*models.User, error) {
	return obj.LoadFamilyMembers()
}
