package primenumber

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
)

// Handler -> get Url has input -> scrapes words of a site
func Handler(w http.ResponseWriter, r *http.Request) {
	var responseJSON []byte
	queryParams := r.URL.Query()
	glog.Info("Prime Number params -> ", queryParams)
	if len(queryParams) == 0 {
		glog.Info("Query Params are missing")
	}
	input, ok := queryParams["input"]
	if !ok {
		glog.Info("URL Param is missing")
	}
	n, err := strconv.Atoi(input[0])
	if err != nil {
		glog.Info(err)
	}

	primeNumbers := make(chan []int)
	go findPrime(n, primeNumbers)
	result := <-primeNumbers

	// Marshal and initialize ResponsJSON
	responseJSON, err = json.Marshal(result)
	if err != nil {
		glog.Info(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJSON))
}

func findPrime(n int, primeNumbers chan []int) {
	glog.Info("Go Routine Started for ", n)
	slice := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			slice = append(slice, i)
		}
	}
	primeNumbers <- slice
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
