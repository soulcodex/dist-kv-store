package storetest

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"codesignal/tests"
)

func TestAcceptanceStoreItemSetSuccess(t *testing.T) {
	container, router := setupStore(t)

	apiKey := gofakeit.BitcoinAddress()
	req := tests.GivenOnePostRequestWithoutAuth("/key", map[string]interface{}{
		"api_key": apiKey,
	})
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusOK)

	stored, _ := container.Services.KeyValueStore.Get("api_key")
	assert.Equal(t, stored, apiKey)
}

func TestAcceptanceStoreItemSetFailsGivenAlreadyExists(t *testing.T) {
	container, router := setupStore(t)
	_ = fillStoreWithElements(container.Services.KeyValueStore, "api_key", 1)

	apiKey := gofakeit.BitcoinAddress()
	req := tests.GivenOnePostRequestWithoutAuth("/key", map[string]interface{}{
		"api_key-1": apiKey,
	})
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusConflict)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"key already exist"}`)
}

func TestAcceptanceStoreItemSetFailsGivenKeyIsEmpty(t *testing.T) {
	_, router := setupStore(t)

	req := tests.GivenOnePostRequestWithoutAuth("/key", map[string]interface{}{
		"": gofakeit.BitcoinAddress(),
	})
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusBadRequest)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"store item key is required"}`)
}

func TestAcceptanceStoreItemSetFailsGivenContentIsEmpty(t *testing.T) {
	_, router := setupStore(t)

	req := tests.GivenOnePostRequestWithoutAuth("/key", map[string]interface{}{
		"api_key": "",
	})
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusBadRequest)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"store item value is required"}`)
}

func TestAcceptanceStoreItemSetFailsGivenMoreThanOneKeyAtTime(t *testing.T) {
	_, router := setupStore(t)

	req := tests.GivenOnePostRequestWithoutAuth("/key", map[string]interface{}{
		"api_key":    gofakeit.BitcoinAddress(),
		"credential": gofakeit.BitcoinAddress(),
	})
	response := tests.WhenExecuteRequestGivenRouter(req, router)
	tests.ThenTheResponseStatusCodeShouldBe(t, response, http.StatusBadRequest)
	tests.ThenTheResponseContentShouldBeEqual(t, response, `{"message":"unexpected content keys"}`)
}
