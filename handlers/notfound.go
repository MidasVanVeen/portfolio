package handlers

import (
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	c := components.NotFound()
	err := components.Layout(c, "Not Found", "").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
