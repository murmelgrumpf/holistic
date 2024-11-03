package handlers

import (
	"holistic/api/templates"
	"holistic/internals/common/assert"
	"holistic/internals/common/routing"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handle(router *chi.Mux) {
	addToRouter(
		append(
			templates.TemplateEndpoints(),
		),
		router)
}

func addToRouter(endpoints []routing.Endpoint, router *chi.Mux) {
	assert.NotNil(router, "Router should not be nil")

	registeredEndpoints := [][]string{
		routing.GET:    make([]string, 0, len(endpoints)),
		routing.POST:   make([]string, 0, len(endpoints)),
		routing.PUT:    make([]string, 0, len(endpoints)),
		routing.DELETE: make([]string, 0, len(endpoints)),
		routing.PATCH:  make([]string, 0, len(endpoints)),
	}
	for _, endpoint := range endpoints {
		var registerFunction func(pattern string, handlerFn http.HandlerFunc)
		switch endpoint.Request {
		case routing.GET:
			registerFunction = router.Get
		case routing.POST:
			registerFunction = router.Post
		case routing.PUT:
			registerFunction = router.Put
		case routing.DELETE:
			registerFunction = router.Delete
		case routing.PATCH:
			registerFunction = router.Patch
		}

		registered := registeredEndpoints[endpoint.Request]

		assert.SliceNotContains(registered, endpoint.Path, "Endpoint already registered", endpoint.Request)
		registerFunction(endpoint.Path, routing.Make(endpoint))
		registered = append(registered, endpoint.Path)

		for _, redirect := range endpoint.Redirects {
			assert.SliceNotContains(registered, redirect, "Endpoint already registered", endpoint.Request)
			registerFunction(redirect, routing.Redirect(endpoint.Path))
			registered = append(registered, redirect)
		}

		registeredEndpoints[endpoint.Request] = registered
	}

}
