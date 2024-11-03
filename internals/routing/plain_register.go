package routing

import (
	"net/http"
)

type EndpointFunc[T any] func(*http.Request) (int, T)

type Endpoint[T any] struct {
	Request   requestType
	Response  responseType
	Path      string
	Function  EndpointFunc[T]
	Redirects []string
}
