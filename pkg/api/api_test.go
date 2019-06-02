package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefineApis(t *testing.T) {
	apiRegistry, err := RegisterAPIs("../../config/config.sample.yaml")
	if err != nil {
		t.Error(err)
	}

	testApi := apiRegistry.APIs["test"]

	assert.Equal(t, "Testing", testApi.Name)
	assert.Equal(t, "test", testApi.Prefix)
	assert.Equal(t, "http://localhost:8080/invalid/url", testApi.UpstreamUrl.String())
	assert.Equal(t, "GET", testApi.Methods[0])
	assert.Equal(t, true, testApi.AllowCrossOrigin)
	assert.Equal(t, true, testApi.Middlewares["logging"])
}
