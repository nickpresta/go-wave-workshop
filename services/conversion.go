package services

import (
	"encoding/json"
	"net/http"

	"github.com/NickPresta/go-wave-workshop/infrastructure"
)

// Conversion represents an instance of an originalAmount in from currency to amount in to currency
type Conversion struct {
	Amount float64 `json:"amount,string,omitempty"`
	From   string  `json:"from,omitempty"`
	To     string  `json:"to,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// Converter is the public service that converts amounts from one currency into another
type Converter struct {
	Infrastructure infrastructure.ConversionInfrastructure
}

// NewConverter creates a new default Converter instance
func NewConverter(apiKey, serviceURL string) *Converter {
	// Passing these are args feels gross
	return &Converter{
		Infrastructure: &infrastructure.DefaultConversionInfrastructure{
			APIKey:     apiKey,
			ServiceURL: serviceURL,
			Client:     &http.Client{},
		},
	}
}

// Convert converts amount in from currency to an amount in to currency.
func (c Converter) Convert(amount float64, from, to string) (*Conversion, error) {
	data, err := c.Infrastructure.Convert(amount, from, to)

	conversion := &Conversion{}
	err = json.Unmarshal(data, conversion)
	if err != nil {
		return nil, err
	}

	return conversion, err
}
