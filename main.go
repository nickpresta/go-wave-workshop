package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/NickPresta/go-wave-workshop/services"
)

const serviceURL = "http://jsonrates.com/convert/"

var apiKey string

func init() {
	apiEnv := os.Getenv("API_KEY")
	if apiEnv == "" {
		apiEnv = "jr-104352fe815d7a92c683578c60fcb877"
	}
	apiKey = apiEnv
}

// RequestPayload is the payload submitted with a conversion request
type RequestPayload struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func main() {
	http.HandleFunc("/convert", ConvertHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

// ConvertHandler handles conversion requests
func ConvertHandler(w http.ResponseWriter, req *http.Request) {
	var body RequestPayload
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	amount, err := strconv.ParseFloat(body.Amount, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	converter := services.NewConverter(apiKey, serviceURL)
	converted, err := converter.Convert(amount, body.From, body.To)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(converted)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
