package server

import (
	"iskaypet-challenge/internal/storage"
	"net/http"
)

type PetHandler struct {
	Storage storage.Repository
}

func (h *PetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.get(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.put(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return

}
