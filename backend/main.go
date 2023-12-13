package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/thalaivar-subu/golang-app/backend/api/crud"
	"github.com/thalaivar-subu/golang-app/backend/api/excel"
	"github.com/thalaivar-subu/golang-app/backend/api/lastday"
	"github.com/thalaivar-subu/golang-app/backend/api/primenumber"
	wordcounter "github.com/thalaivar-subu/golang-app/backend/api/wordcounter"
	database "github.com/thalaivar-subu/golang-app/backend/database"
	"github.com/thalaivar-subu/golang-app/backend/helpers"
	"github.com/thalaivar-subu/golang-app/backend/opentelemetry"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

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
	// Initializing Tracer And Metrics
	opentelemetry.SetupOTelSDK(context.Background(), "golang-app", "1.0")

	helpers.RemoveLogFiles()
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./log/")
	flag.Parse()
	db := database.ConnectMysql()
	defer db.Close()
	router := mux.NewRouter()

	// Auto Instrumentation of Go Mux
	router.Use(otelmux.Middleware("my-server"))

	router.HandleFunc("/", healthCheck)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/wordcounter", helpers.HandlerWrap(wordcounter.Handler)).Methods(http.MethodGet)
	api.HandleFunc("/lastdate", helpers.HandlerWrap(lastday.Handler)).Methods(http.MethodPost)
	api.HandleFunc("/primenumber", helpers.HandlerWrap(primenumber.Handler)).Methods(http.MethodGet)
	api.HandleFunc("/crud", helpers.HandlerWrapWithDb(crud.Handler, db)).Methods("POST", "GET", "DELETE", "PUT")
	api.HandleFunc("/excel", helpers.HandlerWrap(excel.Handler)).Methods(http.MethodGet)
	glog.Info("Server is starting and while listen in " + port)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
	glog.Flush()
}
