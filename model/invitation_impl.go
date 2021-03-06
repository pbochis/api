// This file was automatically generated from
//
//	invitation.go
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

// InvitationKind is the kind used in Datastore to store entities Invitation entities.
const InvitationKind = "Invitation"

// Invitations is just a slice of Invitation.
type Invitations []Invitation

// KeyedInvitation is a struct that embeds Invitation and also contains a Key, mainly used for encoding to JSON.
type KeyedInvitation struct {
	*Invitation
	Key *datastore.Key
}

// Key is a shorthand to fill a KeyedInvitation with an entity and it's key.
func (ƨ *Invitation) Key(key *datastore.Key) *KeyedInvitation {
	return &KeyedInvitation{
		Invitation: ƨ,
		Key:        key,
	}
}

// Key is a shorthand to fill a slice of KeyedInvitation with some entities alongside their keys.
func (ƨ Invitations) Key(keys []*datastore.Key) (keyed []KeyedInvitation) {
	if len(keys) != len(ƨ) {
		panic("Key() called on an slice with len(keys) != len(slice)")
	}

	keyed = make([]KeyedInvitation, len(ƨ))
	for i := range keyed {
		keyed[i] = KeyedInvitation{
			Invitation: &ƨ[i],
			Key:        keys[i],
		}
	}
	return
}

// Put will put this Invitation into Datastore using the given key.
func (ƨ Invitation) Put(ctx context.Context, key *datastore.Key) (*datastore.Key, error) {
	if key != nil {
		return datastore.Put(ctx, key, &ƨ)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Invitation", nil), &ƨ)
}

// PutWithParent can be used to save this Invitation as child of another
// entity.
// This will error if parent == nil.
func (ƨ Invitation) PutWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Invitation", parent), &ƨ)
}

// NewQueryForInvitation prepares a datastore.Query that can be
// used to query entities of type Invitation.
func NewQueryForInvitation() *datastore.Query {
	return datastore.NewQuery("Invitation")
}
