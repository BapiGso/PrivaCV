package main

import (
	"PrivaCV/web"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.FS(web.Assets)))

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
