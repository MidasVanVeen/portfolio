package handlers

import (
	"log/slog"
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := components.About()
	err := components.Layout(c, "About", "/").Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "path", r.URL.Path, "error", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
