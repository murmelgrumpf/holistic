package routing

import (
	"holistic/internals/common/assert"
	"holistic/internals/common/routes"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

var templEndpointsToRegister = make([]Endpoint[templ.Component], 0)

func InitializeEndpoints(router *chi.Mux) {
	addToRouter(templEndpointsToRegister, MakeTempl, router)
}

var registeredEndpoints = [][]string{
	GET:    make([]string, 0),
	POST:   make([]string, 0),
	PUT:    make([]string, 0),
	DELETE: make([]string, 0),
	PATCH:  make([]string, 0),
}

func addToRouter[T any](endpoints []Endpoint[T], make func(endpoint Endpoint[T]) http.HandlerFunc, router *chi.Mux) {
	assert.NotNil(router, "Router should not be nil")

	for _, endpoint := range endpoints {
		var registerFunction func(pattern string, handlerFn http.HandlerFunc)
		switch endpoint.Request {
		case GET:
			registerFunction = router.Get
		case POST:
			registerFunction = router.Post
		case PUT:
			registerFunction = router.Put
		case DELETE:
			registerFunction = router.Delete
		case PATCH:
			registerFunction = router.Patch
		}

		registered := registeredEndpoints[endpoint.Request]

		assert.StringHasPrefix(endpoint.Path, "/", "Path must start with /")
		assert.NotNil(endpoint.Function, "Endpoint function must not be null", endpoint.Request)
		assert.SliceNotContains(registered, endpoint.Path, "Endpoint already registered", endpoint.Request)
		registerFunction(endpoint.Path, make(endpoint))
		registered = append(registered, endpoint.Path)

		for _, redirect := range endpoint.Redirects {
			assert.SliceNotContains(registered, redirect, "Endpoint already registered | Redirect for "+endpoint.Path, endpoint.Request)
			registerFunction(redirect, Redirect(endpoint.Path))
			registered = append(registered, redirect)
			routes.RegisteredRedirects[redirect] = endpoint.Path
		}

		registeredEndpoints[endpoint.Request] = registered
	}

}
