package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const serviceURL = "http://jsonrates.com/convert/"
const apiKey = "jr-104352fe815d7a92c683578c60fcb877"

// Conversion is a convert result
type Conversion struct {
	Amount string `json:"amount,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Error  string `json:"error,omitempty"`
}

// RequestPayload is my payload
type RequestPayload struct {
	Amount string `json:"amount,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}

func main() {
	http.HandleFunc("/convert", ConvertHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

// ConvertHandler is my handler
func ConvertHandler(w http.ResponseWriter, req *http.Request) {
	var body RequestPayload
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	amount, _ := strconv.ParseFloat(body.Amount, 64)
	converted, _ := Convert(body.From, body.To, amount)

	w.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(converted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func doRequest(from, to string, amount float64) (Conversion, error) {
	query := fmt.Sprintf(
		"amount=%f&from=%s&to=%s&apiKey=%s",
		amount,
		from,
		to,
		apiKey,
	)
	resp, _ := http.Get(serviceURL + "?" + query)
	defer resp.Body.Close()

	var conversion Conversion
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&conversion)
	if err != nil {
		return Conversion{}, err
	}
	return conversion, nil
}

// Convert takes an amount and converts it from to to
func Convert(from string, to string, amount float64) (Conversion, error) {
	return doRequest(from, to, amount)
}
