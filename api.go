package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Step - information for the current step
type Step struct {
	Glyph string `json:"glyph"`
}

var seq = [...]string{"|", "/", "-", "\\"}

func handleStep(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Remove the start of the path to get the step number
	step := strings.TrimPrefix(r.URL.Path, "/step/")

	if step != "" {
		num, err := strconv.Atoi(step)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if num < 0 || num >= len(seq) {
			http.Error(w, fmt.Sprintf("Step number %d out of range", num), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(Step{Glyph: seq[num]})
	} else {
		http.Error(w, "No step number given", http.StatusBadRequest)
	}
}

func run() {
	http.HandleFunc("/step/", handleStep)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
