package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func handle_base(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]string{
		"Title":   "日本語 JLPT Analyzer",
		"Heading": "日本語 JLPT Analyzer",
	}
	tmpl.Execute(w, data)
}

func accept_input(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid Form", http.StatusBadRequest)
		return
	}

	input := r.FormValue("langInput")
	log.Println("User submitted: ", input)

	response := fmt.Sprintf("Result: %s", strings.ToUpper(input))
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
