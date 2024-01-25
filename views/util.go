package views

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(component templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Layout(component).Render(r.Context(), w)
	}
}
