package excel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
)

// Handler -> get Url has input -> scrapes words of a site
func Handler(w http.ResponseWriter, r *http.Request) {
	var responseJSON []byte
	input := r.URL.Query().Get("s")
	noOfColumns := r.URL.Query().Get("c")
	col, err := strconv.Atoi(noOfColumns)
	if err != nil {
		glog.Info(err)
	}
	noOfRows := r.URL.Query().Get("r")
	row, err := strconv.Atoi(noOfRows)
	if err != nil {
		glog.Info(err)
	}
	responseJSON, err = json.Marshal(ExcelRowNames(col*row, input))
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}
