package routing

import (
	"holistic/internals/common/assert"
)

type requestType uint8

const (
	GET requestType = iota
	POST
	PUT
	DELETE
	PATCH
)

func (request requestType) String() string {
	switch request {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case PATCH:
		return "PATCH"
	}
	assert.Never("No String variant for requestType", request)
	return ""
}

type responseType uint8

const (
	JSON responseType = iota
	TEMPL
	VOID
)

func (response responseType) String() string {
	switch response {
	case JSON:
		return "JSON"
	case TEMPL:
		return "TEMPL"
	case VOID:
		return "VOID"
	}
	assert.Never("No String variant for responseType", response)
	return ""
}
