package main

import (
	"log"

	"github.com/priyanshu334/tw-bend/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
