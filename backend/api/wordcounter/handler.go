package wordcounter

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/golang/glog"
	"go.opentelemetry.io/otel"
	metric "go.opentelemetry.io/otel/metric"
)

var tracer = otel.Tracer("word-counter-controller")
var meter = otel.Meter("word-counter-controller")

var apiCounter, _ = meter.Int64Counter(
	"subu.counter",
	metric.WithDescription("Number of API calls."),
	metric.WithUnit("{call}"),
)

// Handler -> get Url has input -> scrapes words of a site
func Handler(w http.ResponseWriter, r *http.Request) {
	apiCounter.Add(r.Context(), 1)
	_, span := tracer.Start(r.Context(), "word-counter")
	defer span.End()
	var responseJSON []byte
	url := r.URL.Query().Get("url")
	wordCountMap := make(map[string]int, 0)

	collyIns := colly.NewCollector()
	collyIns.OnHTML("html", func(e *colly.HTMLElement) {
		e.DOM.Find("script,style").Each(func(index int, item *goquery.Selection) {
			item.Remove()
		})
		domContents := strings.Fields(e.DOM.Contents().Text())
		for i := 0; i < len(domContents); i++ {
			if IsAlphaNumeric(domContents[i]) {
				wordCountMap[domContents[i]]++
			}
		}
	})
	collyIns.OnError(func(r *colly.Response, err error) {
		glog.Info(err)
	})
	collyIns.Visit(url)

	responseJSON, err := json.Marshal(wordCountMap)
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}
