package config

import (
	"fmt"
	"net/http"

	"github.com/BernardoDeveloper/powf/src/routes"
	"github.com/go-chi/chi/v5"
)

func StartServer() {
	apiRouter := routes.StartRouter()
	r := chi.NewRouter()
	r.Mount("/api", apiRouter)

	fmt.Println("Starting server on port 3333 🚀")

	http.ListenAndServe(":3333", r)
}
