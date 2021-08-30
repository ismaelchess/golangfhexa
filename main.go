package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ismaelchess/golangfhexa/api"
	"github.com/ismaelchess/golangfhexa/config"
	"github.com/ismaelchess/golangfhexa/domain"
	"github.com/ismaelchess/golangfhexa/repository"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := domain.NewEmployeeService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/employee/{nouser}", handler.Get)
	r.Post("/employee", handler.Post)
	r.Delete("/employee/{nouser}", handler.Delete)
	r.Get("/employee", handler.GetAll)
	r.Put("/employee", handler.Put)
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
