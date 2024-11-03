package routing

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

type RequestType uint8

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
	PATCH
)

type ResponseType uint8

const (
	JSON ResponseType = iota
	TEMPL
	VOID
)

type EndpointFunc func(http.Header, context.Context) (any, error)
type TemplEndpointFunc func(http.Header, context.Context) (templ.Component, error)

type Endpoint struct {
	Request   RequestType
	Response  ResponseType
	Path      string
	Function  EndpointFunc
	Redirects []string
}

func TemplEndpoint(request RequestType, path string, redirects []string, function TemplEndpointFunc) Endpoint {
	return Endpoint{
		Request:   request,
		Response:  TEMPL,
		Path:      path,
		Redirects: redirects,
		Function:  func(header http.Header, context context.Context) (any, error) { return function(header, context) },
	}
}
