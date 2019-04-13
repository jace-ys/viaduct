package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/jace-ys/viaduct/pkg/middleware"
)

type ApiProvider struct {
	server  *http.Server
	negroni *negroni.Negroni // Use negroni to handle middleware
	mux     *mux.Router      // Use gorilla mux internally to do the legwork
}

// Return a new provider with some defaults
func NewApiProvider() *ApiProvider {
	negroni := negroni.New()
	mux := mux.NewRouter().StrictSlash(true)
	return &ApiProvider{
		negroni: negroni,
		mux:     mux,
	}
}

func (p *ApiProvider) Start(port string) error {
	// Return http.Handler to be used as the argument to server.ListenAndServe
	// The mux router needs to be the last item of middleware added to the negroni instance.
	p.negroni.UseHandler(p.mux)

	address := fmt.Sprintf(":%v", port)
	p.server = &http.Server{
		Addr:    address,
		Handler: p.negroni,
	}

	return p.server.ListenAndServe()
}

// Add middleware that will be applied to every request. Middleware handlers are executed in the order they are defined.
func (p *ApiProvider) CommonMiddleware(middlewares ...middleware.Middleware) {
	for _, middleware := range middlewares {
		p.negroni.UseFunc(middleware)
	}
}

func (p *ApiProvider) HandleFunc(methods []string, path string, handlerFunc http.HandlerFunc, middlewares ...middleware.Middleware) {
	// Use the adapter to transform the http.handlerFunc into a http.Handler
	handler := http.HandlerFunc(handlerFunc)

	// Defer to the Handle method
	p.Handle(methods, path, handler, middlewares...)
}

// Register a route for the given path and method. Optionally add middleware.
// see https://github.com/gorilla/mux for options available for the path, including variables.
func (p *ApiProvider) Handle(methods []string, path string, handler http.Handler, middlewares ...middleware.Middleware) {
	// A slice to hold all of the middleware once it's converted (including the handler itself)
	var stack []negroni.Handler

	// The middleware functions have type Middleware but they need to conform to the negroni.Handler interface.
	// By using the negroni.HandlerFunc adapter, they will be given the method required by the interface.
	for _, middleware := range middlewares {
		stack = append(stack, negroni.HandlerFunc(middleware))
	}

	// The handler needs to be treated like middleware.
	// The negroni.Wrap function can convert an http.Handler into a negroni.Handler.
	stack = append(stack, negroni.Wrap(handler))

	// Create the new route using mux
	route := p.mux.NewRoute()

	// If the last character of the path is an asterisk, create a path prefix
	if path[len(path)-1] == '*' {
		// Be sure to strip the asterisk off again
		route.PathPrefix(path[:len(path)-1])
		// Otherwise just add the path normally
	} else {
		route.Path(path)
	}

	// Use a new instance of negroni with our handler stack as the route handler
	route.Handler(negroni.New(stack...))

	// If methods are defined, restrict to those methods only
	if len(methods) > 0 {
		route.Methods(methods...)
	}
}
