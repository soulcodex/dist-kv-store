package storetest

import (
	"net/http"
	"testing"

	"codesignal/tests"
)

func TestAcceptanceStoreItemDeleteSuccess(t *testing.T) {
	container, router := setupStore(t)
	_ = fillStoreWithElements(container.Services.KeyValueStore, "api_key", 1)

	req := tests.GivenOneDeleteRequestWithoutAuth("/key/api_key-1")
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusOK)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"key deleted successfully"}`)
}

func TestAcceptanceStoreItemDeleteFailsGivenKeyDoesNotExists(t *testing.T) {
	container, router := setupStore(t)
	_ = fillStoreWithElements(container.Services.KeyValueStore, "api_key", 1)

	req := tests.GivenOneDeleteRequestWithoutAuth("/key/api_key")
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusNotFound)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"key not found"}`)
}
