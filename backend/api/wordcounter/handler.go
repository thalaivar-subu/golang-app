package wordcounter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/golang/glog"
)

// Handler -> get Url has input -> scrapes words of a site
func Handler(w http.ResponseWriter, r *http.Request) {
	var responseJSON []byte
	queryParams := r.URL.Query()
	glog.Info("Wordounter params -> ", queryParams)
	if len(queryParams) == 0 {
		glog.Info("Query Params are missing")
	}
	url, ok := queryParams["url"]
	if !ok {
		glog.Info("URL Param is missing")
	}
	response, err := http.Get(url[0])
	if err != nil {
		glog.Fatal(err)
	}
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			glog.Fatal(err)
		}
		response.Body.Close()

		// Stringify Response Body
		htmlBody := string(bodyBytes)

		// Remove Script Tags
		regexToMatchScriptTags := regexp.MustCompile(`<script\b[^>]*>([\s\S]*?)<\/script>`)
		textWithOutJS := regexToMatchScriptTags.ReplaceAllString(htmlBody, "")

		// Remove Tags from Html
		regexToMatchTags := regexp.MustCompile(`<[^>]*>`)
		textWithOutHTML := regexToMatchTags.ReplaceAllString(textWithOutJS, "")

		// Match Words from string
		regexToGetWords := regexp.MustCompile(`(\b[a-zA-Z][a-zA-Z]+|\b[A-Z]\b)`)
		wordsInHTML := regexToGetWords.FindAllString(textWithOutHTML, -1)

		// Find the count of each word
		wordCountMap := make(map[string]int)
		for i := 0; i < len(wordsInHTML); i++ {
			eachWord := wordsInHTML[i]
			wordCountMap[eachWord]++
		}
		// Marshal and initialize ResponsJSON
		responseJSON, err = json.Marshal(wordCountMap)
		if err != nil {
			glog.Info(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}
