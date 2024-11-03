package templates

import (
	"holistic/api/templates/globalmap"
	"holistic/api/templates/test"
	"holistic/internals/common/routing"
)

func TemplateEndpoints() []routing.Endpoint {
	return []routing.Endpoint{
		globalmap.GlobalMapGet,
		test.TestGet,
	}
}
