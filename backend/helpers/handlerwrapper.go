package helpers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/thalaivar-subu/golang-app/backend/types"
)

func HandlerWrapWithDb(h types.HandlerWithDb, db *gorm.DB) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		EnableCors(&w)
		defer func() {
			if r := recover(); r != nil {
				glog.Info(r)
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		h(w, r, db)
	}
}

func HandlerWrap(h types.Handler) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		EnableCors(&w)
		defer func() {
			if r := recover(); r != nil {
				glog.Info(r)
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		h(w, r)
	}
}

func requestLogger(r *http.Request) {
	method := r.Method
	uri := r.URL.String()
	glog.Info("method: "+method, " uri: "+uri, " startTime: "+time.Now().String())
}
