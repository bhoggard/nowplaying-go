package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", index)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		fmt.Println("There was an error parsing the template:", err.Error())
	}
	tmpl_vars := make(map[string]interface{})
	tmpl_vars["counterstream"] = counterstream()
	t.Execute(w, tmpl_vars)
}
