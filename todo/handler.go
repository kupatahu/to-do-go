package todo

import (
	"encoding/json"
	"net/http"

	"github.com/kupatahu/to-do-go/entity"
)

type Usecase interface {
	Get() ([]*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(todo *entity.Todo) (*entity.Todo, error)
}

type handler struct {
	usecase Usecase
}

func newHandler(usecase Usecase) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	todo, err := h.usecase.Get()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todoJson, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(todoJson)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var todo *entity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err = h.usecase.Create(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todoJson, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(todoJson)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var todo *entity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err = h.usecase.Update(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todoJson, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(todoJson)
}
