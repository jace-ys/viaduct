package logging

import (
	"net/http"
	"strings"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/utils/format"
)

type serviceContext struct {
	Name       string
	Host       string
	Method     string
	RequestURI string
}

func getServiceContext(r *http.Request, registry *config.ServiceRegistry) *serviceContext {
	// Trim prefix to obtain actual request URI
	for _, service := range registry.Services {
		prefix := format.AddSlashes(service.Prefix)
		actualURI := strings.TrimPrefix(r.RequestURI, prefix)

		// Find the service that matches the request prefix
		if strings.Contains(r.RequestURI, prefix) {
			return &serviceContext{
				Name:       service.Name,
				Host:       service.UpstreamUrl.Host,
				Method:     r.Method,
				RequestURI: format.SingleJoiningSlash(service.UpstreamUrl.Path, actualURI),
			}
		}
	}

	// Else return `Unknown Endpoint`
	return &serviceContext{
		Name:       "Unknown Endpoint",
		Host:       r.Host,
		Method:     r.Method,
		RequestURI: r.RequestURI,
	}
}
