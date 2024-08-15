package storetest

import (
	"net/http"
	"testing"

	"codesignal/tests"
)

func TestAcceptanceStoreItemGetSuccess(t *testing.T) {
	container, router := setupStore(t)
	err := fillStoreWithElements(container.Services.KeyValueStore, "api_key", 1)
	if err != nil {
		t.Fatal(err)
	}

	req := tests.GivenOneGetRequestWithoutAuth("/key/api_key-1")
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusOK)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"api_key-1":"<<PRESENCE>>"}`)
}

func TestAcceptanceStoreItemGetNotFound(t *testing.T) {
	_, router := setupStore(t)

	req := tests.GivenOneGetRequestWithoutAuth("/key/api_key-1")
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusNotFound)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"key not found"}`)
}
