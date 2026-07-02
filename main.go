package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Payments = NewPaymentsHandler()

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!\n"))
	})

	r.Post("/payment", handlePost)
	r.Get("/payment/{id}", handleGet)
	r.Get("/payment", handleGetAll)
	r.Delete("/payment/{id}", handleDelete)
	r.Put("/payment/{id}", handlePut)

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {

	payments := Payments.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func handleGet(w http.ResponseWriter, r *http.Request) {

	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := Payments.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)

}

func handlePut(w http.ResponseWriter, r *http.Request) {

	id, err := parseId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload Payment
	err = json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedPayment, err := Payments.Update(id, payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedPayment)
	}
}

func parseId(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = Payments.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var payload Payment

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Payments.Add(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(payload)
	}
}
