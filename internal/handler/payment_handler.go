package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	. "payment_api/internal/model"
	. "payment_api/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {

	var pCreate PaymentCreate

	err := json.NewDecoder(r.Body).Decode(&pCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.service.Create(&pCreate)

	if errors.Is(err, ErrAlreadyExists) {
		http.Error(w, ErrAlreadyExists.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.Delete(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func parseId(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, err
	}
	return int64(id), nil
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {

	id, err := parseId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// var pUpdate Payment
	var pUpdate PaymentUpdate
	err = json.NewDecoder(r.Body).Decode(&pUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.service.Update(id, &pUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payment)
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {

	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.service.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)

}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {

	payments, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}
