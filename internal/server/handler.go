package server

import (
	"encoding/json"
	"iskaypet-challenge/internal/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReadRepository interface {
	Get(id int) (*domain.Pet, error)
	List() ([]*domain.Pet, error)
}

type PetRHandler struct {
	storage ReadRepository
}

func NewPetReaderHandler(repo ReadRepository) *PetRHandler {
	return &PetRHandler{storage: repo}
}

type WriteRepository interface {
	Insert(p *domain.Pet) error
}

type PetWHandler struct {
	storage WriteRepository
}

func NewPetWriterHandler(repo WriteRepository) *PetWHandler {
	return &PetWHandler{storage: repo}
}

func (h *PetRHandler) List(w http.ResponseWriter, r *http.Request) {
	pets, err := h.storage.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pets)
}

func (h *PetRHandler) Get(w http.ResponseWriter, r *http.Request) {
	i := mux.Vars(r)["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet, err := h.storage.Get(id)
	if err != nil {
		if err.Error() == "no content" {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)
}

func (h *PetWHandler) Create(w http.ResponseWriter, r *http.Request) {
	var pet domain.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pet.IsValidDate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.storage.Insert(&pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
