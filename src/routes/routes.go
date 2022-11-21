package routes

import (
	getpath "github.com/BernardoDeveloper/powf/src/useCases/GetPath"
	userfile "github.com/BernardoDeveloper/powf/src/useCases/UserFile"
	"github.com/go-chi/chi/v5"
)

func StartRouter() *chi.Mux {
	ApiRouter := chi.NewRouter()

	ApiRouter.Post("/path", getpath.Path)
	ApiRouter.Post("/upload", userfile.SaveFile)

	return ApiRouter
}
