package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func createHandler() Handler {
	u := newUsecase()
	return newHandler(u)
}

func RegisterRoutes(r chi.Router) {
	h := createHandler()

	r.Route("/ping", func(r chi.Router) {
		r.Get("/", h.Get)
	})
}
