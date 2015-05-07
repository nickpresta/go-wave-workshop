package infrastructure

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// ConversionInfrastructure is an interface for making calls to a third-party for conversion
type ConversionInfrastructure interface {
	Convert(amount float64, from, to string) ([]byte, error)
}

// DefaultConversionInfrastructure is responsible for interfacing with our third party
type DefaultConversionInfrastructure struct {
	APIKey     string
	ServiceURL string
	Client     *http.Client
}

// Convert converts amount from a from currency to a to currency
func (i DefaultConversionInfrastructure) Convert(amount float64, from, to string) ([]byte, error) {
	parsedURL, _ := url.Parse(i.ServiceURL)

	originalAmount := strconv.FormatFloat(amount, 'f', -1, 64)
	query := url.Values{}
	query.Set("amount", originalAmount)
	query.Set("from", from)
	query.Set("to", to)
	query.Set("apiKey", i.APIKey)
	parsedURL.RawQuery = query.Encode()

	resp, err := i.Client.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
