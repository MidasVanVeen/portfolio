package handlers

import (
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	c := components.Index()
	err := components.Layout(c, "Home").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
