package main

import (
	"github.com/gorilla/mux"
	"handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/publish", handlers.PublishHandler).Methods("POST").Schemes("http")
	r.HandleFunc("/v1/consume", handlers.ConsumerHandler).Methods("POST").Schemes("http")
	http.ListenAndServe(":8080", r)
}


