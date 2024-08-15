package tests

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"codesignal/cmd/di"
	"codesignal/internal/pkg/utils"
)

type HttpTestRouterFactory func(di *di.OktaDistributedKeyValueStorageContainer) *httprouter.Router

func GivenOneGetRequestWithoutAuth(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func GivenOnePostRequestWithoutAuth(url string, raw map[string]interface{}) *http.Request {
	body, _ := utils.MarshalFromMap(raw)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func GivenOneDeleteRequestWithoutAuth(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func WhenExecuteRequestGivenRouter(req *http.Request, router *httprouter.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func ThenTheResponseStatusCodeShouldBe(t *testing.T, response *httptest.ResponseRecorder, expected int) {
	if response.Code != expected {
		t.Errorf("Expected response code %d. Got %d\n", expected, response.Code)
	}
}

func ThenTheResponseContentShouldBeEqual(t *testing.T, response *httptest.ResponseRecorder, expected string) {
	jsonAssert := jsonassert.New(t)

	if response.Body.String() == "" {
		assert.Equal(t, expected, response.Body.String())
		return
	}

	jsonAssert.Assertf(response.Body.String(), expected)

	if t.Failed() {
		t.Errorf("Expected response body %s. Got %s\n", expected, response.Body.String())
	}
}
