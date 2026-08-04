package main

import (
	"fmt"
	"net/http"

	"payment_api/internal/handler"
	"payment_api/internal/repository"
	"payment_api/internal/service"
	"payment_api/pkg/config"
	"payment_api/pkg/db"
	"payment_api/pkg/logger"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log, err := logger.New(cfg.Logger.Level)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	database, err := db.New(cfg.Database)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	repo := repository.NewPaymentPg(database)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, log)

	r := chi.NewRouter()
	router := handler.NewRouterWrapper(h, r)
	router.Map()

	fmt.Println("Listening on", cfg.Server.Addr)
	err = http.ListenAndServe(cfg.Server.Addr, r)
	if err != nil {
		panic(err)
	}
}
