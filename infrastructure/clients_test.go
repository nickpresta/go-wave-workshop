package infrastructure_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NickPresta/go-wave-workshop/infrastructure"
)

func TestConvert(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check for payload data shape in here
		fmt.Fprintf(w, "OK")
	}))
	defer server.Close()

	sut := infrastructure.DefaultConversionInfrastructure{
		APIKey:     "123",
		ServiceURL: server.URL,
		Client:     &http.Client{},
	}
	bytes, err := sut.Convert(float64(100), "USD", "CAD")
	if err != nil {
		t.Error(err)
	}

	if len(bytes) <= 0 {
		t.Errorf("Expected more than 0 bytes")
	}
}
