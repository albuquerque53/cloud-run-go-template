package main

import (
	"gcloudrungo/internal/application"
	"log"
)

func main() {
	log.Print("starting server...")

	application.Run()
}
