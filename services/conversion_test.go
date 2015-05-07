package services_test

import (
	"fmt"
	"testing"

	"github.com/NickPresta/go-wave-workshop/services"
)

type MockInfrastructure struct {
	amount float64
	from   string
	to     string
}

func (m *MockInfrastructure) setResponse(amount float64, from, to string) {
	m.amount = amount
	m.from = from
	m.to = to
}

func (m MockInfrastructure) Convert(amount float64, from, to string) ([]byte, error) {
	return []byte(fmt.Sprintf(`{"from":"%s","to":"%s","amount":"%f"}`, m.from, m.to, m.amount)), nil
}

func TestConvert(t *testing.T) {
	converter := &services.Converter{
		Infrastructure: &MockInfrastructure{},
	}

	tests := []struct {
		amount float64
		from   string
		to     string
	}{
		{float64(123.123), "USD", "CAD"},
		{float64(48.39), "EUR", "GBP"},
	}

	for _, test := range tests {
		// Seed each test run
		converter.Infrastructure.(*MockInfrastructure).setResponse(test.amount, test.from, test.to)

		actual, err := converter.Convert(test.amount, test.from, test.to)
		if err != nil {
			t.Error(err)
		}
		expected := &services.Conversion{
			Amount: test.amount,
			From:   test.from,
			To:     test.to,
		}
		if actual.Amount != test.amount || actual.From != test.from || actual.To != test.to {
			t.Errorf("Convert(%v,%v,%v) = %+v; want %+v", test.amount, test.from, test.to, actual, expected)
		}
	}
}
