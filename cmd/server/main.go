package main

import (
	"gonvue/pkg/config"
	"gonvue/pkg/http/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	config := config.New()

	server := &http.Server{
		Handler:      router.New(config),
		Addr:         fmt.Sprintf("127.0.0.1:%s", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server started on port :%s\n", config.Port)
	log.Fatal(server.ListenAndServe())
}
