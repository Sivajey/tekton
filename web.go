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
	log.Print("Google Search home page")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
