package lastday

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

type body struct {
	Date string
}

// Handler -> get date has input -> returns last day of the month
func Handler(w http.ResponseWriter, r *http.Request) {
	var responseJSON []byte
	decoder := json.NewDecoder(r.Body)

	var data body
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	dateMap := map[string]interface{}{
		"Date":           data.Date,
		"LastDayOfMonth": LastDate(data.Date)}
	responseJSON, err = json.Marshal(dateMap)
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}

// curl -XPOST http://127.0.0.1:3001/api/v1/lastdate -d '{"date":"2020-02-18T01:50"}'
