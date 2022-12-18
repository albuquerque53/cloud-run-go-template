package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")

	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
}
