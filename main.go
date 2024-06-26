package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/public", PublicHandler)
	http.HandleFunc("/refresh", RefreshHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
