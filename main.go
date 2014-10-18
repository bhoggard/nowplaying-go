package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/second-inversion", secondInversionJSON)
	http.HandleFunc("/counterstream", counterstreamJSON)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":"+port, nil)
}
