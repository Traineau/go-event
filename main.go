package main

import (
	"goevent/router"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	log.Print("\nServer started on port " + port)

	err = http.ListenAndServe(":"+port, newRouter)
	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
