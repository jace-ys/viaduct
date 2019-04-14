package logging

import (
	"net/http"
	"strings"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/utils/format"
)

type apiContext struct {
	Name       string
	Host       string
	Method     string
	RequestURI string
}

func getApiContext(r *http.Request, registry *config.ApiRegistry) *apiContext {
	// Trim prefix to obtain actual request URI
	for _, apiDefinition := range registry.Apis {
		prefix := format.AddSlashes(apiDefinition.Prefix)
		actualURI := strings.TrimPrefix(r.RequestURI, prefix)

		// Find the API that matches the request prefix
		if strings.Contains(r.RequestURI, prefix) {
			return &apiContext{
				Name:       apiDefinition.Name,
				Host:       apiDefinition.UpstreamUrl.Host,
				Method:     r.Method,
				RequestURI: format.SingleJoiningSlash(apiDefinition.UpstreamUrl.Path, actualURI),
			}
		}
	}

	// Else return `Unknown Endpoint`
	return &apiContext{
		Name:       "Unknown Endpoint",
		Host:       r.Host,
		Method:     r.Method,
		RequestURI: r.RequestURI,
	}
}
