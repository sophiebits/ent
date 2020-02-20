// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package models

import (
	"context"
	"sync"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/cast"
	"github.com/lolopinto/ent/ent/privacy"
	"github.com/lolopinto/ent/ent/sql"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models/configs"
)

const (
	// ContactType is the node type for the Contact object. Used to identify this node in edges and other places.
	ContactType ent.NodeType = "contact"

	// ContactToAllowListEdge is the edgeType for the contact to allowlist edge.
	ContactToAllowListEdge ent.EdgeType = "f6ecacb9-1d4f-47bb-8f18-f7d544450ea2"
)

// Contact represents the `Contact` model
type Contact struct {
	ent.Node
	privacy.AlwaysDenyPrivacyPolicy
	EmailAddress  string   `db:"email_address"`
	FirstName     string   `db:"first_name"`
	LastName      string   `db:"last_name"`
	UserID        string   `db:"user_id"`
	Favorite      *bool    `db:"favorite" graphql:"_"`
	NumberOfCalls *int     `db:"number_of_calls" graphql:"_"`
	Pi            *float64 `db:"pi" graphql:"_"`
	Viewer        viewer.ViewerContext
}

type Contacts map[string]*Contact

// ContactResult stores the result of loading a Contact. It's a tuple type which has 2 fields:
// a Contact and an error
type ContactResult struct {
	Contact *Contact
	Err     error
}

func (res *ContactResult) Error() string {
	return res.Err.Error()
}

// ContactsResult stores the result of loading a slice of Contacts. It's a tuple type which has 2 fields:
// a []*Contact and an error
type ContactsResult struct {
	Contacts []*Contact
	Err      error
}

func (res *ContactsResult) Error() string {
	return res.Err.Error()
}

// contactLoader is an ent.PrivacyBackedLoader which is used to
// load Contact
type contactLoader struct {
	nodes   map[string]*Contact
	errs    map[string]error
	results []*Contact
	v       viewer.ViewerContext
	m       sync.Mutex
}

func (res *contactLoader) GetNewInstance() ent.DBObject {
	var contact Contact
	contact.Viewer = res.v
	return &contact
}

func (res *contactLoader) GetConfig() ent.Config {
	return &configs.ContactConfig{}
}

func (res *contactLoader) SetPrivacyResult(id string, obj ent.DBObject, err error) {
	res.m.Lock()
	defer res.m.Unlock()
	if err != nil {
		res.errs[id] = err
	} else if obj != nil {
		// TODO kill results?
		ent := obj.(*Contact)
		res.nodes[id] = ent
		res.results = append(res.results, ent)
	}
}

func (res *contactLoader) GetEntForID(id string) *Contact {
	return res.nodes[id]
}

// hmm make private...
func (res *contactLoader) List() []*Contact {
	return res.results
}

func (res *contactLoader) getFirstInstance() *Contact {
	if len(res.results) == 0 {
		return nil
	}
	return res.results[0]
}

func (res *contactLoader) getFirstErr() error {
	for _, err := range res.errs {
		return err
	}
	return nil
}

// NewContactLoader returns a new contactLoader which is used to load one or more Contacts
func NewContactLoader(v viewer.ViewerContext) *contactLoader {
	return &contactLoader{
		nodes: make(map[string]*Contact),
		errs:  make(map[string]error),
		v:     v,
	}
}

// IsNode is needed by gqlgen to indicate that this implements the Node interface in GraphQL
func (contact Contact) IsNode() {}

// GetType returns the NodeType of this entity. In this case: ContactType
func (contact *Contact) GetType() ent.NodeType {
	return ContactType
}

// GetViewer returns the viewer for this entity.
func (contact *Contact) GetViewer() viewer.ViewerContext {
	return contact.Viewer
}

// LoadContactFromContext loads the given Contact given the context and id
func LoadContactFromContext(ctx context.Context, id string) (*Contact, error) {
	v, err := viewer.ForContext(ctx)
	if err != nil {
		return nil, err
	}
	return LoadContact(v, id)
}

// GenLoadContactFromContext loads the given Contact given the context and id
func GenLoadContactFromContext(ctx context.Context, id string) <-chan *ContactResult {
	res := make(chan *ContactResult)
	go func() {
		v, err := viewer.ForContext(ctx)
		if err != nil {
			res <- &ContactResult{
				Err: err,
			}
			return
		}
		res <- <-(GenLoadContact(v, id))
	}()
	return res
}

// LoadContact loads the given Contact given the viewer and id
func LoadContact(v viewer.ViewerContext, id string) (*Contact, error) {
	loader := NewContactLoader(v)
	err := ent.LoadNode(v, id, loader)
	return loader.nodes[id], err
}

// GenLoadContact loads the given Contact given the id
func GenLoadContact(v viewer.ViewerContext, id string) <-chan *ContactResult {
	res := make(chan *ContactResult)
	go func() {
		var result ContactResult
		loader := NewContactLoader(v)
		result.Err = <-ent.GenLoadNode(v, id, loader)
		result.Contact = loader.nodes[id]
		res <- &result
	}()
	return res
}

// LoadContacts loads multiple Contacts given the ids
func LoadContacts(v viewer.ViewerContext, ids ...string) ([]*Contact, error) {
	loader := NewContactLoader(v)
	err := ent.LoadNodes(v, ids, loader)
	return loader.results, err
}

// GenLoadContacts loads multiple Contacts given the ids
func GenLoadContacts(v viewer.ViewerContext, ids ...string) <-chan *ContactsResult {
	res := make(chan *ContactsResult)
	go func() {
		loader := NewContactLoader(v)
		var result ContactsResult
		result.Err = <-ent.GenLoadNodes(v, ids, loader)
		result.Contacts = loader.results
		res <- &result
	}()
	return res
}

func LoadContactIDFromEmailAddress(emailAddress string) (string, error) {
	loader := NewContactLoader(viewer.LoggedOutViewer())
	data, err := ent.LoadNodeRawDataViaQueryClause(
		loader,
		sql.Eq("email_address", emailAddress),
	)
	if err != nil {
		return "", err
	}
	return cast.ToUUIDString(data["id"])
}

func LoadContactFromEmailAddress(v viewer.ViewerContext, emailAddress string) (*Contact, error) {
	loader := NewContactLoader(v)
	err := ent.LoadNodeViaQueryClause(v, loader, sql.Eq("email_address", emailAddress))
	if err != nil {
		return nil, err
	}
	return loader.getFirstInstance(), loader.getFirstErr()
}

// GenUser returns the User associated with the Contact instance
func (contact *Contact) GenUser() <-chan *UserResult {
	return GenLoadUser(contact.Viewer, contact.UserID)
}

// LoadUser returns the User associated with the Contact instance
func (contact *Contact) LoadUser() (*User, error) {
	return LoadUser(contact.Viewer, contact.UserID)
}

// GenContactEmails returns the ContactEmails associated with the Contact instance
func (contact *Contact) GenContactEmails() <-chan *ContactEmailsResult {
	res := make(chan *ContactEmailsResult)
	go func() {
		loader := NewContactEmailLoader(contact.Viewer)
		var result ContactEmailsResult
		result.Err = <-ent.GenLoadNodesViaQueryClause(contact.Viewer, loader, sql.Eq("contact_id", contact.ID))
		result.ContactEmails = loader.results
		res <- &result
	}()
	return res
}

// LoadContactEmails returns the ContactEmails associated with the Contact instance
func (contact *Contact) LoadContactEmails() ([]*ContactEmail, error) {
	loader := NewContactEmailLoader(contact.Viewer)
	err := ent.LoadNodesViaQueryClause(contact.Viewer, loader, sql.Eq("contact_id", contact.ID))
	return loader.results, err
}

// LoadAllowListEdges returns the AllowList edges associated with the Contact instance
func (contact *Contact) LoadAllowListEdges() ([]*ent.AssocEdge, error) {
	return ent.LoadEdgesByType(contact.ID, ContactToAllowListEdge)
}

// GenAllowListEdges returns the User edges associated with the Contact instance
func (contact *Contact) GenAllowListEdges() <-chan *ent.AssocEdgesResult {
	return ent.GenLoadEdgesByType(contact.ID, ContactToAllowListEdge)
}

// GenAllowList returns the Users associated with the Contact instance
func (contact *Contact) GenAllowList() <-chan *UsersResult {
	res := make(chan *UsersResult)
	go func() {
		loader := NewUserLoader(contact.Viewer)
		var result UsersResult
		result.Err = <-ent.GenLoadNodesByType(contact.Viewer, contact.ID, ContactToAllowListEdge, loader)
		result.Users = loader.results
		res <- &result
	}()
	return res
}

// LoadAllowList returns the Users associated with the Contact instance
func (contact *Contact) LoadAllowList() ([]*User, error) {
	loader := NewUserLoader(contact.Viewer)
	err := ent.LoadNodesByType(contact.Viewer, contact.ID, ContactToAllowListEdge, loader)
	return loader.results, err
}

// LoadAllowListEdgeFor loads the ent.AssocEdge between the current node and the given id2 for the AllowList edge.
func (contact *Contact) LoadAllowListEdgeFor(id2 string) (*ent.AssocEdge, error) {
	return ent.LoadEdgeByType(contact.ID, id2, ContactToAllowListEdge)
}

// GenAllowListEdgeFor provides a concurrent API to load the ent.AssocEdge between the current node and the given id2 for the AllowList edge.
func (contact *Contact) GenLoadAllowListEdgeFor(id2 string) <-chan *ent.AssocEdgeResult {
	return ent.GenLoadEdgeByType(contact.ID, id2, ContactToAllowListEdge)
}

// DBFields is used by the ent framework to load the ent from the underlying database
func (contact *Contact) DBFields() ent.DBFields {
	return ent.DBFields{
		"id": func(v interface{}) error {
			var err error
			contact.ID, err = cast.ToUUIDString(v)
			return err
		},
		"created_at": func(v interface{}) error {
			var err error
			contact.CreatedAt, err = cast.ToTime(v)
			return err
		},
		"updated_at": func(v interface{}) error {
			var err error
			contact.UpdatedAt, err = cast.ToTime(v)
			return err
		},
		"email_address": func(v interface{}) error {
			var err error
			contact.EmailAddress, err = cast.ToString(v)
			return err
		},
		"first_name": func(v interface{}) error {
			var err error
			contact.FirstName, err = cast.ToString(v)
			return err
		},
		"last_name": func(v interface{}) error {
			var err error
			contact.LastName, err = cast.ToString(v)
			return err
		},
		"user_id": func(v interface{}) error {
			var err error
			contact.UserID, err = cast.ToString(v)
			return err
		},
		"favorite": func(v interface{}) error {
			var err error
			contact.Favorite, err = cast.ToNullableBool(v)
			return err
		},
		"number_of_calls": func(v interface{}) error {
			var err error
			contact.NumberOfCalls, err = cast.ToNullableInt(v)
			return err
		},
		"pi": func(v interface{}) error {
			var err error
			contact.Pi, err = cast.ToNullableFloat(v)
			return err
		},
	}
}

var _ ent.Entity = &Contact{}
