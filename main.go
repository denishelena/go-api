package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func HandleHello(w http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Hello %s", vars["name"]),
	})
}

func HandlePerson(w http.ResponseWriter, rq *http.Request) {
	var p Person

	err := json.NewDecoder(rq.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Nome: %s - Idade: %d\n", p.Name, p.Age)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HandleHello).Methods("GET")
	r.HandleFunc("/person", HandlePerson).Methods("POST")

	http.ListenAndServe(":8080", r)
}
