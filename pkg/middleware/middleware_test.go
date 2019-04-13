package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureError(t *testing.T) {
	middlewares := map[string]bool{
		"testing": false,
		"logging": true,
	}

	middlewareStack, err := Configure(middlewares, Registry{})

	assert.Equal(t, 0, len(middlewareStack))
	assert.Equal(t, "Unknown middleware declared in config: logging", err.Error())
}
