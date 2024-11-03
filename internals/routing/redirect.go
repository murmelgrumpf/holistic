package routing

import (
	"net/http"
	"strings"
)

func Redirect(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectPath := path

		prevQuery := r.URL.RawQuery
		if len(prevQuery) > 0 {
			if strings.Contains(path, "?") {
				redirectPath += "&"
			} else {
				redirectPath += "?"
			}
			redirectPath += r.URL.RawQuery
		}

		http.Redirect(w, r, redirectPath, http.StatusSeeOther)
	}
}
