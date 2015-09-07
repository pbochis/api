package tests

import (
	"fmt"
	"net/http/httptest"
	"net/mail"
	"os"
	"testing"

	"golang.org/x/net/context"

	"github.com/coduno/api/controllers"
	"github.com/coduno/api/model"
	"github.com/coduno/api/tests/testUtils"

	"google.golang.org/appengine/aetest"
)

var ctx context.Context
var instance aetest.Instance

func TestMain(m *testing.M) {
	var err error
	var callback func()
	ctx, callback, err = aetest.NewContext()
	if err != nil {
		os.Exit(1)
	}
	defer callback()

	instance, err = aetest.NewInstance(nil)
	if err != nil {
		os.Exit(1)
	}
	testUtils.MockData(ctx)
	os.Exit(m.Run())
}

// TestCompanyController tests stuff
func TestPostCompany(t *testing.T) {
	company := model.Company{
		Address: mail.Address{
			Name:    "name",
			Address: "email@google.com",
		},
	}

	req, err := instance.NewRequest("POST", "/companies", testUtils.RequestBody(t, company))

	if err != nil {
		fmt.Printf("Err: %v", err)
		t.Fatal(err)

	}

	w := httptest.NewRecorder()

	status, err := controllers.PostCompany(ctx, w, req)
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Fatal(status)
	}
}
