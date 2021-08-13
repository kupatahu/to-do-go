package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/kupatahu/to-do-go/healthcheck"
	"github.com/kupatahu/to-do-go/todo"
)

func main() {
	r := chi.NewRouter()

	healthcheck.RegisterRoutes(r)
	todo.RegisterRoutes(r)

	addr := ":" + os.Getenv("PORT")
	fmt.Println("server started on" + addr)
	err := http.ListenAndServe(addr, r)

	if err != nil {
		log.Fatal(err)
	}
}
