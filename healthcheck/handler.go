package healthcheck

import (
	"net/http"
)

type Usecase interface {
	Get() string
}

type handler struct {
	usecase Usecase
}

func newHandler(usecase Usecase) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	v := h.usecase.Get()

	w.Write([]byte(v + "\n"))
}

func Get2(usecase Usecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := usecase.Get()
	
		w.Write([]byte(v + "\n"))
	}
}
