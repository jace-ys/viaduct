package middleware

import (
	"fmt"
	"net/http"
)

// This replicates the negroni.HandlerFunc type but abstracts the code from the library
type Middleware func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// Registry associates each Middleware with its name, to be used by Configure
type Registry map[string]Middleware

// Spits out a stack of middlewares using the provided registry, based on the middleware config for each API
func Configure(middlewares map[string]bool, registry Registry) (stack []Middleware, e error) {
	// Add middlewares declared in config to stack
	for key, value := range middlewares {
		if value {
			// Return an error if middleware does not exist in registry
			if registry[key] == nil {
				err := fmt.Errorf("Unknown middleware declared in config: %s", key)
				return stack, err
			}
			stack = append(stack, registry[key])
		}
	}
	// Return the middleware stack
	return stack, nil
}
