package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	teachers := []string{
		"Siyovush",
		"Nurullo",
		"Fariz",
	}

	students := []string{
		"Alijon",
		"Mehrdod",
		"Fozil",
		"Muhammad",
	}

	teachersHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		switch r.Method {
		case "GET":
		case "DELETE":
			io.WriteString(w, "Method not allowed")
		}
		io.WriteString(w, fmt.Sprint(teachers))
	}

	studentsHandler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, fmt.Sprint(students))
	}

	// our endpoints:
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, `Some cool html`)
	})
	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/groups", groupsHandler)

	log.Println("Server is running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func groupsHandler(w http.ResponseWriter, _ *http.Request) {
	groups := []string{
		"C#",
		"Go",
		"Golang",
		"Golangjon",
	}

	resp, _ := json.Marshal(groups)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, string(resp))
}
