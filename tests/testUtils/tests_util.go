package testUtils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/mail"
	"testing"
	"time"

	"github.com/coduno/api/model"
	"github.com/coduno/api/util/password"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

var MockedCompanyKey *datastore.Key
var MockedCompanyUserKey *datastore.Key
var MockedUserKey *datastore.Key
var MockedTaskKey *datastore.Key
var MockedChallengeKey *datastore.Key

func MockData(ctx context.Context) {
	var err error
	MockedCompanyKey, err = model.Company{
		Address: mail.Address{
			Name:    "Coduno",
			Address: "team@cod.uno",
		},
	}.Put(ctx, nil)
	if err != nil {
		panic(err)
	}

	pw, _ := password.Hash([]byte("passwordpassword"))
	MockedCompanyUserKey, err = model.User{
		Address: mail.Address{
			Name:    "John",
			Address: "john@cod.uno",
		},
		Nick:           "John",
		HashedPassword: pw,
		Company:        MockedCompanyKey,
	}.Put(ctx, nil)
	if err != nil {
		panic(err)
	}

	MockedUserKey, err = model.User{
		Address: mail.Address{
			Name:    "Andy",
			Address: "andy@gmail.com",
		},
		Nick:           "andy",
		HashedPassword: pw,
	}.Put(ctx, nil)
	if err != nil {
		panic(err)
	}

	MockedTaskKey, err = model.Task{
		Assignment: model.Assignment{
			Name:         "Task one",
			Description:  "Description of task one",
			Instructions: "Instructions of task one",
			Duration:     time.Hour,
			Endpoints: model.Endpoints{
				WebInterface: "coding-task",
			},
		},
		Languages:    []string{"py", "java"},
		SkillWeights: model.SkillWeights{0.1, 0.4, 0.5},
	}.Put(ctx, nil)
	if err != nil {
		panic(err)
	}

	MockedChallengeKey, err = model.Challenge{
		Assignment: model.Assignment{
			Name:         "Challenge one",
			Description:  "Description of challenge one",
			Instructions: "Instructions of challenge one yay",
			Duration:     time.Hour,
			Endpoints: model.Endpoints{
				WebInterface: "sequential-challenge",
			},
		},
		Tasks: []*datastore.Key{MockedTaskKey},
	}.PutWithParent(ctx, MockedCompanyKey)
	if err != nil {
		panic(err)
	}
}

// CreateAndSaveCompany creates a new Company and saves it to datastore using the provided context
func CreateAndSaveCompany(ctx context.Context, name, email string) (model.Company, *datastore.Key, error) {
	if name == "" {
		name = "Default Company"
	}
	if email == "" {
		email = "office@defaul.com"
	}
	company := model.Company{
		Address: mail.Address{
			Name:    name,
			Address: email,
		},
	}
	var err error
	var key *datastore.Key

	key, err = company.Put(ctx, nil)
	if err != nil {
		return model.Company{}, nil, err
	}
	return company, key, nil
}

func RequestBody(t *testing.T, data interface{}) io.Reader {
	var jsonData, err = json.Marshal(data)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return bytes.NewReader(jsonData)
}
