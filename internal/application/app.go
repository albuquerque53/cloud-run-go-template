package application

import (
	"gcloudrungo/internal/application/handler"
	"log"
	"net/http"
	"os"
)

func Run() {
	http.HandleFunc("/", handler.Entrypoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
