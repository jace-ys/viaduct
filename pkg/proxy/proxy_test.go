package proxy

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jace-ys/viaduct/pkg/config"
)

type typicodePost struct {
	UserId int    `json:userId`
	Id     int    `json:id`
	Title  string `json:title`
	Body   string `json:body`
}

func TestProxy(t *testing.T) {
	apiRegistry, err := config.RegisterApis("../../config/config.sample.yaml")
	if err != nil {
		t.Error(err)
	}

	testApi := apiRegistry.Apis["typicode"]
	proxy := New(&testApi)

	proxyHandler := StripPrefix(proxy.api.Prefix, proxy)

	req, _ := http.NewRequest("GET", "/typicode/posts/1", nil)
	res := httptest.NewRecorder()
	proxyHandler.ServeHTTP(res, req)

	assert.Equal(t, "JSONPlaceholder", proxy.api.Name)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "application/json; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.Equal(t, true, res.Body.Len() > 0)

	// Test content of JSON body
	jsonBody := &typicodePost{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(jsonBody)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, jsonBody.Id)
	assert.Equal(t, true, len(jsonBody.Title) > 0)
}
