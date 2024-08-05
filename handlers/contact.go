package handlers

import (
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	c := components.Contact()
	err := components.Layout(c, "Contact", "/contact").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
