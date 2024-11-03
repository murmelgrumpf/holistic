package routing

import (
	"holistic/internals/common/logger"
	"net/http"

	"github.com/a-h/templ"
)

func MakeTempl(endpoint Endpoint[templ.Component]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {

				logger.Error("PANIC", "err", err)

				if IsDocument(r.Header) {
					http.Redirect(w, r, "/?toast=INTERNAL_ERROR", http.StatusSeeOther)
					return
				}
				if IsHtmx(r.Header) {
					status, component := InternalError()
					WriteTempl(w, r, status, component)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		status, component := endpoint.Function(r)
		WriteTempl(w, r, status, component)
	}
}

func WriteTempl(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := component.Render(r.Context(), w); err != nil {

		logger.Error("Internal Error ocurred", err)

		_, component = InternalError()
		if err := component.Render(r.Context(), w); err != nil {
			logger.Panic("InternalError Component threw an Error ;)")
		}
	}

}
