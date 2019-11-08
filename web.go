package main

import (
	"io"
	"net/http"

	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		panic(err)
	}
}

func main() {
	log.Print("Google search is ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
