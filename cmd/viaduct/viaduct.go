package viaduct

import (
	"net/http"

	"github.com/jace-ys/viaduct/pkg/api"
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

	// Load API definitions declared in config file
	apiRegistry, err := api.RegisterAPIs(conf.ConfigFile)
	if err != nil {
		return err
	}

	// Create new ApiProvider
	apiProvider := router.NewApiProvider()

	// Register middlewares to be used by middleware.Configure
	middlewareRegistry := middleware.Registry{
		"logging": logging.CreateMiddleware(log.Request(), &apiRegistry),
	}

	// Apply middleware common to every request route
	apiProvider.CommonMiddleware(middlewareRegistry["logging"])

	for _, apiDefinition := range apiRegistry.APIs {
		// Create a new reverse proxy for each API
		reverseProxy := proxy.New(&apiDefinition)

		// Add surrounding slashes to the prefix
		prefix := format.AddSlashes(apiDefinition.Prefix)

		// Strip prefix before passing to proxy
		handler := http.StripPrefix(prefix, reverseProxy)

		// Configure a stack of middlewares for each API
		middlewareStack, err := middleware.Configure(apiDefinition.Middlewares, middlewareRegistry)
		if err != nil {
			return err
		}

		log.Debug().Printf(
			format.Apis,
			apiDefinition.Name,
			prefix,
			apiDefinition.UpstreamUrl,
			format.Methods(apiDefinition.Methods),
			apiDefinition.AllowCrossOrigin,
			format.Middleware(apiDefinition.Middlewares),
		)

		// Handle requests with signature /prefix/* using the proxy and apply the stack of middlewares to be handled
		apiProvider.Handle(apiDefinition.Methods, prefix+"*", handler, middlewareStack...)
	}

	// Start ListenAndServe using apiProvider.server
	log.Debug().Println("Server listening on port", conf.Port)
	return apiProvider.Start(conf.Port)
}
