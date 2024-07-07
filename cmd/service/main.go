package main

import (
	"log"

	"cinematic.com/sso/internal/infrastructure/container"
)

func main() {
	c, err := container.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
