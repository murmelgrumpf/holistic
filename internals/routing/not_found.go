package routing

import (
	"net/http"
)

func HandleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsDocument(r.Header) {

			redirectPath := "/?toast=NOT_FOUND"

			prevQuery := r.URL.RawQuery
			if len(prevQuery) > 0 {
				redirectPath += "&" + r.URL.RawQuery
			}

			http.Redirect(w, r, redirectPath, http.StatusSeeOther)
			return
		}
		if IsHtmx(r.Header) {
			status, component := NotFound()
			WriteTempl(w, r, status, component)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}
}
