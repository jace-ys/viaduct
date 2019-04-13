package router

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jace-ys/viaduct/pkg/middleware"
)

var response = "Invalid URL accessed!"

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(response))
}

func TestNoMiddleware(t *testing.T) {
	middlewareStack, err := middleware.Configure(map[string]bool{}, middleware.Registry{})
	if err != nil {
		t.Error(err)
	}

	apiProvider := NewApiProvider()
	apiProvider.HandleFunc([]string{"GET"}, "/test/*", testHandler, middlewareStack...)

	handlers := apiProvider.negroni.Handlers()

	assert.Equal(t, 0, len(handlers))
}

func TestApiProviderSuccess(t *testing.T) {
	apiProvider := NewApiProvider()
	apiProvider.HandleFunc([]string{"GET"}, "/test/*", testHandler)

	server := httptest.NewServer(apiProvider.mux)
	defer server.Close()

	endpoint := server.URL + "/test/success"
	req, _ := http.NewRequest("GET", endpoint, nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
	assert.Equal(t, true, len(body) > 0)
	assert.Equal(t, true, bytes.Contains(body, []byte(response)))
}

func TestApiProviderError(t *testing.T) {
	apiProvider := NewApiProvider()
	apiProvider.HandleFunc([]string{"POST"}, "/test/*", testHandler)

	server := httptest.NewServer(apiProvider.mux)
	defer server.Close()

	endpoint := server.URL + "/test/error"
	req, _ := http.NewRequest("GET", endpoint, nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.Equal(t, true, len(body) == 0)
}
