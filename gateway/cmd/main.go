package main

import (
	"github.com/mytord/fs/gateway/internal"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	backendUri, err := url.Parse(os.Getenv("BACKEND_URI"))

	if err != nil || backendUri.Host == "" {
		panic("invalid BACKEND_URI setting")
	}

	secret, ok := os.LookupEnv("JWT_SECRET")

	if !ok || secret == "" {
		panic("invalid JWT_SECRET setting")
	}

	gateway := internal.NewApiGateway(backendUri, secret)

	mux := http.NewServeMux()

	mux.HandleFunc("/", gateway.Handle)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Authorization"},
		ExposedHeaders: []string{"X-Set-Token", "X-Token-Expires"},
	}).Handler(mux)

	log.Print("gateway server started")

	srv := &http.Server{
		Addr:         ":80",
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
