package main

import (
	"anomaly-detection-go/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Alive)
	router.HandleFunc("/anomaly-detection", handler.AnomalyDetection)
	log.Fatal(http.ListenAndServe(":8080", router))
}
