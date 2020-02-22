// Code generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

package models

import (
	"context"
	"sync"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/cast"
	"github.com/lolopinto/ent/ent/privacy"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/internal/test_schema/models/configs"
)

const (
	// AddressType is the node type for the Address object. Used to identify this node in edges and other places.
	AddressType ent.NodeType = "address"
)

// Address represents the `Address` model
type Address struct {
	ent.Node
	privacy.AlwaysDenyPrivacyPolicy
	City          string   `db:"city"`
	Country       string   `db:"country"`
	ResidentNames []string `db:"resident_names"`
	State         string   `db:"state"`
	StreetAddress string   `db:"street_address"`
	Zip           string   `db:"zip"`
	Viewer        viewer.ViewerContext
}

type Addresses map[string]*Address

// AddressResult stores the result of loading a Address. It's a tuple type which has 2 fields:
// a Address and an error
type AddressResult struct {
	Address *Address
	Err     error
}

func (res *AddressResult) Error() string {
	return res.Err.Error()
}

// AddressesResult stores the result of loading a slice of Addresss. It's a tuple type which has 2 fields:
// a []*Address and an error
type AddressesResult struct {
	Addresses []*Address
	Err       error
}

func (res *AddressesResult) Error() string {
	return res.Err.Error()
}

// addressLoader is an ent.PrivacyBackedLoader which is used to
// load Address
type addressLoader struct {
	nodes   map[string]*Address
	errs    map[string]error
	results []*Address
	v       viewer.ViewerContext
	m       sync.Mutex
}

func (res *addressLoader) GetNewInstance() ent.DBObject {
	return res.GetNewAddress()
}

func (res *addressLoader) GetNewAddress() *Address {
	var address Address
	address.Viewer = res.v
	return &address
}

func (res *addressLoader) GetConfig() ent.Config {
	return &configs.AddressConfig{}
}

func (res *addressLoader) SetPrivacyResult(id string, obj ent.DBObject, err error) {
	res.m.Lock()
	defer res.m.Unlock()
	if err != nil {
		res.errs[id] = err
	} else if obj != nil {
		// TODO kill results?
		ent := obj.(*Address)
		res.nodes[id] = ent
		res.results = append(res.results, ent)
	}
}

func (res *addressLoader) GetEntForID(id string) *Address {
	return res.nodes[id]
}

// hmm make private...
func (res *addressLoader) List() []*Address {
	return res.results
}

func (res *addressLoader) getFirstInstance() *Address {
	if len(res.results) == 0 {
		return nil
	}
	return res.results[0]
}

func (res *addressLoader) getFirstErr() error {
	for _, err := range res.errs {
		return err
	}
	return nil
}

// NewAddressLoader returns a new addressLoader which is used to load one or more Addresses
func NewAddressLoader(v viewer.ViewerContext) *addressLoader {
	return &addressLoader{
		nodes: make(map[string]*Address),
		errs:  make(map[string]error),
		v:     v,
	}
}

// IsNode is needed by gqlgen to indicate that this implements the Node interface in GraphQL
func (address Address) IsNode() {}

// GetType returns the NodeType of this entity. In this case: ContactType
func (address *Address) GetType() ent.NodeType {
	return AddressType
}

// GetViewer returns the viewer for this entity.
func (address *Address) GetViewer() viewer.ViewerContext {
	return address.Viewer
}

// LoadAddressFromContext loads the given Address given the context and id
func LoadAddressFromContext(ctx context.Context, id string) (*Address, error) {
	v, err := viewer.ForContext(ctx)
	if err != nil {
		return nil, err
	}
	return LoadAddress(v, id)
}

// GenLoadAddressFromContext loads the given Address given the context and id
func GenLoadAddressFromContext(ctx context.Context, id string) <-chan *AddressResult {
	res := make(chan *AddressResult)
	go func() {
		address, err := LoadAddressFromContext(ctx, id)
		res <- &AddressResult{
			Err:     err,
			Address: address,
		}
	}()
	return res
}

// LoadAddress loads the given Address given the viewer and id
func LoadAddress(v viewer.ViewerContext, id string) (*Address, error) {
	loader := NewAddressLoader(v)
	err := ent.LoadNode(v, id, loader)
	return loader.nodes[id], err
}

// GenLoadAddress loads the given Address given the id
func GenLoadAddress(v viewer.ViewerContext, id string) <-chan *AddressResult {
	res := make(chan *AddressResult)
	go func() {
		address, err := LoadAddress(v, id)
		res <- &AddressResult{
			Err:     err,
			Address: address,
		}
	}()
	return res
}

// LoadAddresses loads multiple Addresses given the ids
func LoadAddresses(v viewer.ViewerContext, ids ...string) ([]*Address, error) {
	loader := NewAddressLoader(v)
	err := ent.LoadNodes(v, ids, loader)
	return loader.results, err
}

// GenLoadAddresses loads multiple Addresses given the ids
func GenLoadAddresses(v viewer.ViewerContext, ids ...string) <-chan *AddressesResult {
	res := make(chan *AddressesResult)
	go func() {
		results, err := LoadAddresses(v, ids...)
		res <- &AddressesResult{
			Err:       err,
			Addresses: results,
		}
	}()
	return res
}

// DBFields is used by the ent framework to load the ent from the underlying database
func (address *Address) DBFields() ent.DBFields {
	return ent.DBFields{
		"id": func(v interface{}) error {
			var err error
			address.ID, err = cast.ToUUIDString(v)
			return err
		},
		"created_at": func(v interface{}) error {
			var err error
			address.CreatedAt, err = cast.ToTime(v)
			return err
		},
		"updated_at": func(v interface{}) error {
			var err error
			address.UpdatedAt, err = cast.ToTime(v)
			return err
		},
		"city": func(v interface{}) error {
			var err error
			address.City, err = cast.ToString(v)
			return err
		},
		"country": func(v interface{}) error {
			var err error
			address.Country, err = cast.ToString(v)
			return err
		},
		"resident_names": func(v interface{}) error {
			return cast.UnmarshallJSON(v, &address.ResidentNames)
		},
		"state": func(v interface{}) error {
			var err error
			address.State, err = cast.ToString(v)
			return err
		},
		"street_address": func(v interface{}) error {
			var err error
			address.StreetAddress, err = cast.ToString(v)
			return err
		},
		"zip": func(v interface{}) error {
			var err error
			address.Zip, err = cast.ToString(v)
			return err
		},
	}
}

// UnsupportedScan flags that we can't call StructScan() on the ent to get data out of the db, have to always use MapScan() and DBFields() method above
func (address *Address) UnsupportedScan() bool {
	return true
}

var _ ent.Entity = &Address{}
