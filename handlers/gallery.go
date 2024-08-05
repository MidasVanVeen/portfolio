package handlers

import (
	"net/http"

	components "github.com/midasvanveen/portfolio/v2/components"
	"github.com/midasvanveen/portfolio/v2/db"
)

type GalleryHandler struct {
	store *db.ProjectStore
}

func NewGalleryHandler(store *db.ProjectStore) *GalleryHandler {
	return &GalleryHandler{
		store: store,
	}
}

func (h *GalleryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projects, err := h.store.GetAllProjects()
	if err != nil {
		http.Error(w, "Error fetching from db", http.StatusInternalServerError)
		return
	}

	c := components.Gallery(projects)
	err = components.Layout(c, "Gallery", "/gallery").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
