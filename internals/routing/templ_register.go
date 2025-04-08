package routing

import (
	"holistic/api/templates/structure"
	"holistic/internals/common/assert"
	"holistic/internals/common/option"
	"holistic/internals/common/templating"
	"net/http"

	"github.com/a-h/templ"
)

type TemplEndpointFunc EndpointFunc[templ.Component]

type PartialEndpoint struct {
	Request   requestType
	Path      string
	Title     option.Option[string]
	Redirects []string
	Function  TemplEndpointFunc
}

type PageOnlyEndpoint struct {
	Request   requestType
	Path      string
	Title     string
	Redirects []string
	Function  TemplEndpointFunc
}

type PageEndpointFunc func(*http.Request, bool) (int, templ.Component)

type PageEndpoint struct {
	Request   requestType
	Path      string
	Title     string
	Redirects []string
	Function  PageEndpointFunc
}

func (endpoint PartialEndpoint) Register() {

	assert.StringHasPrefix(endpoint.Path, "/", "Path must start with /")
	assert.NotNil(endpoint.Function, "Endpoint function must not be null", endpoint.Request)

	var partialRedirects = make([]string, len(endpoint.Redirects))
	for i, url := range endpoint.Redirects {
		partialRedirects[i] = "/partial" + url
	}

	templEndpointsToRegister = append(templEndpointsToRegister, Endpoint[templ.Component]{
		Request:   endpoint.Request,
		Path:      "/partial" + endpoint.Path,
		Redirects: partialRedirects,
		Function: func(r *http.Request) (int, templ.Component) {
			status, component := endpoint.Function(r)
			if endpoint.Title.IsNone() {
				return status, component
			}
			return status, templating.Wrap(structure.Title(endpoint.Title.Unwrap()), component)
		},
	})
}

func (endpoint PageOnlyEndpoint) Register() {

	assert.StringHasPrefix(endpoint.Path, "/", "Path must start with /")
	assert.NotNil(endpoint.Function, "Endpoint function must not be null", endpoint.Request)

	templEndpointsToRegister = append(templEndpointsToRegister, Endpoint[templ.Component]{
		Request:   endpoint.Request,
		Response:  TEMPL,
		Path:      endpoint.Path,
		Redirects: endpoint.Redirects,
		Function: func(r *http.Request) (int, templ.Component) {
			status, component := endpoint.Function(r)
			return status, templating.Wrap(structure.Base(endpoint.Title, r.URL.Query().Get("toast")), component)
		},
	})
}

func (endpoint PageEndpoint) Register() {
	PageOnlyEndpoint{
		Request:   endpoint.Request,
		Path:      endpoint.Path,
		Title:     endpoint.Title,
		Redirects: endpoint.Redirects,
		Function: func(r *http.Request) (int, templ.Component) {
			return endpoint.Function(r, false)
		},
	}.Register()

	PartialEndpoint{
		Request:   endpoint.Request,
		Path:      endpoint.Path,
		Title:     option.Some(endpoint.Title),
		Redirects: endpoint.Redirects,
		Function: func(r *http.Request) (int, templ.Component) {
			return endpoint.Function(r, true)
		},
	}.Register()
}
