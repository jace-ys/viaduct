package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// This replicates the negroni.HandlerFunc type but abstracts the code from the library
type Middleware func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// Registry associates each Middleware with its name, to be used by Configure
type Registry map[string]Middleware

// Spits out a stack of middlewares using the provided registry, based on the middleware config for each service
func Configure(definition map[string]bool, registry Registry) (stack []Middleware, e error) {
	// Add middlewares declared in config to stack
	for key, value := range definition {
		if value {
			// Return an error if middleware does not exist in registry
			if registry[key] == nil {
				err := fmt.Errorf("Unknown middleware declared in config: %s", strings.Title(key))
				return stack, err
			}
			stack = append(stack, registry[key])
		}
	}
	// Return the middleware stack
	return stack, nil
}

func Format(m map[string]bool) string {
	var s string

	if len(m) == 0 {
		return "None"
	}

	for key, val := range m {
		s = s + fmt.Sprintf("\n		%s: %t", strings.Title(key), val)
	}

	return s
}
