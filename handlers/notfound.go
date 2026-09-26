package handlers

import (
	"log/slog"
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	c := components.NotFound()
	err := components.Layout(c, "Not Found", "").Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "path", r.URL.Path, "error", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
