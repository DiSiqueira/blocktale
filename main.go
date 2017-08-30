package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
	"net/http"
	"github.com/disiqueira/blocktale/handler/tale"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Error.Fatal(err)
	}

	dbMap, err := db.NewDB(cfg.MySQLDSN, 5, 5, 0)
	if err != nil {
		log.Error.Fatal(err)


		r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/tale", func(r chi.Router) {
		r.Get("/{taleID}", tale.NewGetTaleHandler().ServeHTTP)
		r.Post("/{taleID}", tale.NewPostTaleHandler().ServeHTTP)
	})
}