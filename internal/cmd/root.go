package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/api"
	"github.com/litvinovmitch11/avito-merch-store/internal/server"
	"github.com/rs/zerolog"
)

func Run() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	server := server.NewServer(&logger)

	r := chi.NewMux()
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "localhost:8080",
	}

	log.Fatal(s.ListenAndServe())
}
