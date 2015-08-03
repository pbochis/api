// This file was automatically generated from
//
//	junit_submission.go
//
// by
//
//	generator
//
// DO NOT EDIT

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Save will put this JunitSubmission into Datastore using the given key.
func (ƨ JunitSubmission) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &ƨ)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "junitSubmissions", nil), &ƨ)
}

// SaveWithParent can be used to save this JunitSubmission as child of another
// entity.
// This will error if parent == nil.
func (ƨ JunitSubmission) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "junitSubmissions", parent), &ƨ)
}

// NewQueryForJunitSubmission prepares a datastore.Query that can be
// used to query entities of type JunitSubmission.
func NewQueryForJunitSubmission() *datastore.Query {
	return datastore.NewQuery("junitSubmissions")
}

type JunitSubmissionHandler struct{}

func (ƨ JunitSubmissionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if r.URL.Path == "" {
		var results JunitSubmissions
		keys, _ := NewQueryForJunitSubmission().GetAll(ctx, &results)
		json.NewEncoder(w).Encode(results.Key(keys))
		return
	}

	k, _ := datastore.DecodeKey(r.URL.Path)
	var entity JunitSubmission
	datastore.Get(ctx, k, &entity)
	json.NewEncoder(w).Encode(entity)
}

func ServeJunitSubmission(prefix string, muxes ...*http.ServeMux) {
	path := prefix + "junitSubmissions" + "/"

	if len(muxes) == 0 {
		http.Handle(path, http.StripPrefix(path, JunitSubmissionHandler{}))
	}

	for _, mux := range muxes {
		mux.Handle(path, http.StripPrefix(path, JunitSubmissionHandler{}))
	}
}
