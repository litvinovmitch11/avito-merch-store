package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}

	log.Fatal(s.ListenAndServe())
}
