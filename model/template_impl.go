// This file was automatically generated from
//
//	template.go
//
// by
//
//	generator
//
// DO NOT EDIT

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// TemplateKind is the kind used in Datastore to store entities Template entities.
const TemplateKind = "Template"

// Templates is just a slice of Template.
type Templates []Template

// KeyedTemplate is a struct that embeds Template and also contains a Key, mainly used for encoding to JSON.
type KeyedTemplate struct {
	*Template
	Key *datastore.Key
}

// Key is a shorthand to fill a KeyedTemplate with an entity and it's key.
func (ƨ *Template) Key(key *datastore.Key) *KeyedTemplate {
	return &KeyedTemplate{
		Template: ƨ,
		Key:      key,
	}
}

// Key is a shorthand to fill a slice of KeyedTemplate with some entities alongside their keys.
func (ƨ Templates) Key(keys []*datastore.Key) (keyed []KeyedTemplate) {
	if len(keys) != len(ƨ) {
		panic("Key() called on an slice with len(keys) != len(slice)")
	}

	keyed = make([]KeyedTemplate, len(ƨ))
	for i := range keyed {
		keyed[i] = KeyedTemplate{
			Template: &ƨ[i],
			Key:      keys[i],
		}
	}
	return
}

// Put will put this Template into Datastore using the given key.
func (ƨ Template) Put(ctx context.Context, key *datastore.Key) (*datastore.Key, error) {
	if key != nil {
		return datastore.Put(ctx, key, &ƨ)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Template", nil), &ƨ)
}

// PutWithParent can be used to save this Template as child of another
// entity.
// This will error if parent == nil.
func (ƨ Template) PutWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Template", parent), &ƨ)
}

// NewQueryForTemplate prepares a datastore.Query that can be
// used to query entities of type Template.
func NewQueryForTemplate() *datastore.Query {
	return datastore.NewQuery("Template")
}
