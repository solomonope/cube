package worker

import chi "github.com/go-chi/chi/v5"

type Api struct {
	Address string
	Port    int64
	Worker  *Worker
	Router  chi.Mux
}
