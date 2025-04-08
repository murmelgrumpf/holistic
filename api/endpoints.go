package api

import (
	_ "holistic/api/templates/attributions"
	_ "holistic/api/templates/globalmap"
	_ "holistic/api/templates/test"
	"holistic/internals/routing"

	"github.com/go-chi/chi/v5"
)

func InitializeEndpoints(router *chi.Mux) {
	routing.InitializeEndpoints(router)
}
