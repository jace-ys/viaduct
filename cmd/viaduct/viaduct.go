package viaduct

import (
	"net/http"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/middleware"
	"github.com/jace-ys/viaduct/pkg/middleware/logging"
	"github.com/jace-ys/viaduct/pkg/proxy"
	"github.com/jace-ys/viaduct/pkg/router"
	"github.com/jace-ys/viaduct/pkg/utils/format"
	"github.com/jace-ys/viaduct/pkg/utils/log"
)

func Start() error {
	// Load environmental variables
	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Load service definitions declared in config file
	serviceRegistry, err := config.RegisterServices(conf.ConfigFile)
	if err != nil {
		return err
	}

	// Create new ApiProvider
	apiProvider := router.NewApiProvider()

	// Register middlewares to be used by middleware.Configure
	middlewareRegistry := middleware.Registry{
		"logging": logging.CreateMiddleware(log.Request(), &serviceRegistry),
	}

	// // Apply middleware common to every request route
	// apiProvider.CommonMiddleware(middlewareRegistry["logging"])

	for _, serviceDefinition := range serviceRegistry.Services {
		// Create a new reverse proxy for each service
		reverseProxy := proxy.New(&serviceDefinition)

		// Add surrounding slashes to the prefix
		prefix := format.AddSlashes(serviceDefinition.Prefix)

		// Strip prefix before passing to proxy
		handler := http.StripPrefix(prefix, reverseProxy)

		// Configure a stack of middlewares for each service
		middlewareStack, err := middleware.Configure(serviceDefinition.Middlewares, middlewareRegistry)
		if err != nil {
			return err
		}

		log.Debug().Printf(
			format.Services,
			serviceDefinition.Name,
			prefix,
			serviceDefinition.UpstreamUrl,
			format.Methods(serviceDefinition.Methods),
			serviceDefinition.AllowCrossOrigin,
			format.Middleware(serviceDefinition.Middlewares),
		)

		// Handle requests with signature /prefix/* using the proxy and apply the stack of middlewares to be handled
		apiProvider.Handle(serviceDefinition.Methods, prefix+"*", handler, middlewareStack...)
	}

	// Start ListenAndServe using apiProvider.server
	log.Debug().Println("Server listening on port", conf.Port)
	return apiProvider.Start(conf.Port)
}
