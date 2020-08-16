package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/thalaivar-subu/golang-app/backend/api/lastday"
	"github.com/thalaivar-subu/golang-app/backend/api/primenumber"
	wordcounter "github.com/thalaivar-subu/golang-app/backend/api/wordcounter"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

const port = ":3001"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	glog.Info("healtcheck -> server is up")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": 200, "message": "Server is Up :) "}`))
}

func main() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./log/")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", healthCheck)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/wordcounter", wordcounter.Handler).Methods(http.MethodGet)
	api.HandleFunc("/lastdate", lastday.Handler).Methods(http.MethodPost)
	api.HandleFunc("/primenumber", primenumber.Handler).Methods(http.MethodGet)

	glog.Info("Server is starting and while listen in " + port)
	log.Fatal(http.ListenAndServe(port, router))

	glog.Flush()
}
