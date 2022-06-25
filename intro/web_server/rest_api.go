package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet/{name}", greet).Methods(http.MethodGet)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var name = vars["name"]
	_, err := w.Write([]byte("Hello " + name + "!"))
	if err != nil {
		return
	}
}
