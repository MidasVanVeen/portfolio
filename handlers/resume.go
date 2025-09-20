package handlers

import (
	"net/http"

	"slices"
	components "github.com/midasvanveen/portfolio/v2/components"
	"github.com/midasvanveen/portfolio/v2/db"
)

type ResumeHandler struct {
	store *db.ResumeStore
}

func NewResumeHandler(store *db.ResumeStore) *ResumeHandler {
	return &ResumeHandler{
		store: store,
	}
}

func (h *ResumeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	entries, err := h.store.GetAllResumeEntries()
	slices.Reverse(entries)
	if err != nil {
		http.Error(w, "Error fetching from db", http.StatusInternalServerError)
		return
	}

	c := components.Resume(entries)
	err = components.Layout(c, "Resume", "/resume").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
