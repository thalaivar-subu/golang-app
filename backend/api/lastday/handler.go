package lastday

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	t, err := time.Parse("2006-01-02", strings.Split(data.Date, "T")[0])

	if err != nil {
		fmt.Println(err)
	}
	year := t.Year()
	month := t.Month()
	lastDate := time.Date(year, month+1, 0, 0, 0, 0, 0, t.Location()).Day()
	dateMap := make(map[string]interface{})
	dateMap["Date"] = data.Date
	dateMap["LastDayOfMonth"] = lastDate
	glog.Info("last day APi response ", dateMap)
	responseJSON, err = json.Marshal(dateMap)
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}

// curl -XPOST http://127.0.0.1:3001/api/v1/lastdate -d '{"date":"2020-02-18T01:50"}'
