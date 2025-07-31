package main

import (
	"fmt"
	"log"
	"net/http"
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
	var stats StatTracker
	if stats.WordList == nil {
		stats.WordList = make(map[string]int)
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid Form", http.StatusBadRequest)
	}

	input := r.FormValue("langInput")
	lang_tokens, err := tagInput(input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusInternalServerError)
	}
	htmxString, _ := stats.buildHtmxOut(lang_tokens)
	log.Println("User submitted: ", input)

	response := fmt.Sprintf("Result: %s", htmxString)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
