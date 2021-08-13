package todo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kupatahu/to-do-go/db"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func createHandler() Handler {
	db := db.Open()
	repo := newRepo(db)
	usecase := newUsecase(repo)
	return newHandler(usecase)
}

func RegisterRoutes(r chi.Router) {
	h := createHandler()

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Post("/", h.Create)
		r.Put("/", h.Update)
	})
}
