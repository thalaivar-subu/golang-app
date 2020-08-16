package lastday

import (
	"strings"
	"time"

	"github.com/golang/glog"
)

func LastDate(date string) int {
	t, err := time.Parse("2006-01-02", strings.Split(date, "T")[0])
	if err != nil {
		glog.Info(err)
	}
	year := t.Year()
	month := t.Month()
	lastDate := time.Date(year, month+1, 0, 0, 0, 0, 0, t.Location()).Day()
	return lastDate
}
