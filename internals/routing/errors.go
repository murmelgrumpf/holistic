package routing

import (
	"holistic/api/templates/toast"
	"net/http"

	"github.com/a-h/templ"
)

type ApiErrorFunc func() (Status int, Error templ.Component)

var InternalError ApiErrorFunc = func() (Status int, Error templ.Component) {
	return http.StatusInternalServerError, toast.Error("The action could not be performed.")
}

var NotFound ApiErrorFunc = func() (Status int, Error templ.Component) {
	return http.StatusNotFound, toast.Error("The requested action could not be found.")
}
