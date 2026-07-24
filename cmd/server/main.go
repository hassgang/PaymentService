package main

import (
	"fmt"
	"net/http"
	"os"

	. "payment_api/internal/handler"
	. "payment_api/internal/repository"
	. "payment_api/internal/service"

	"github.com/go-chi/chi/v5"
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
	routerWrapper := NewRouterWrapper(handler, r)
	routerWrapper.Map()

	fmt.Println("Listening...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}
