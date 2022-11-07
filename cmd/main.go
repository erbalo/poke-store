package main

import (
	"fmt"
	"poke-store/internal/adapters/http"
)

func main() {
	fmt.Printf("Initializing project")
	server := http.NewServer()
	server.Start()
}
