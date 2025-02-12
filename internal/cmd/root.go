package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	api "github.com/litvinovmitch11/avito-merch-store/internal/generated"
	"github.com/litvinovmitch11/avito-merch-store/internal/server"
)

func Run() {
	server := server.NewServer()

	r := chi.NewMux()
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
