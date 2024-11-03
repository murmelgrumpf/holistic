package elements

import (
	"holistic/api/elements/navigation"

	"github.com/a-h/templ"
)

var (
	NAVIGATION_MAP = register("navigation_map", navigation.NavigationMap)
)

type componentFunction func() templ.Component

type elementType struct {
	id        string
	component templ.Component
}

func register(id string, component componentFunction) elementType {
	return elementType{
		id:        id,
		component: component(),
	}
}
