package templating

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Wrap(wrapper templ.Component, toWrap templ.Component) templ.Component {
    return templ.ComponentFunc( func (ctx context.Context, w io.Writer) error {
        ctx = templ.WithChildren(ctx, toWrap)
        return wrapper.Render(ctx, w)
    })
}


