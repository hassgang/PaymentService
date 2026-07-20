package main

import (
	"fmt"
	"net/http"
	"os"

	. "payment_api/internal/handler"
	. "payment_api/internal/repository"
	. "payment_api/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connection := os.Getenv("CONNECTION_STRING")

	fmt.Println("Connection str is", connection)

	db, err := sqlx.Connect("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	repo := NewPaymentPg(db)
	service := NewService(repo)
	handler := NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!\n"))
	})

	r.Post("/payment", handler.HandlePost)
	r.Get("/payment/{id}", handler.HandleGet)
	r.Get("/payment", handler.HandleGetAll)
	r.Delete("/payment/{id}", handler.HandleDelete)
	r.Put("/payment/{id}", handler.HandleUpdate)

	fmt.Println("Listening...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
