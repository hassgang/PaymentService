package handler

import (
	"github.com/go-chi/chi/v5"
)

func NewRouterWrapper(handler *Handler, r *chi.Mux) *RouterWrapper {
	return &RouterWrapper{handler: handler, r: r}
}

type RouterWrapper struct {
	handler *Handler
	r       *chi.Mux
}

func (w *RouterWrapper) Map() {
	w.r.Post("/payment", w.handler.HandlePost)
	w.r.Get("/payment", w.handler.HandleGetAll)
	w.r.Get("/payment/{id}", w.handler.HandleGet)
	w.r.Delete("/payment/{id}", w.handler.HandleDelete)
	w.r.Put("/payment/{id}", w.handler.HandleUpdate)
}
