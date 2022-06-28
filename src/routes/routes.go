package routes

import (
	getpath "github.com/BernardoDeveloper/powf/src/useCases/GetPath"
	"github.com/go-chi/chi/v5"
)

func StartRouter() *chi.Mux {
	ApiRouter := chi.NewRouter()

	ApiRouter.Post("/path", getpath.Path)

	return ApiRouter
}
