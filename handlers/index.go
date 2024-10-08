package handlers

import (
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := components.About()
	err := components.Layout(c, "About", "/").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
