package belajargojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("customer_encode.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Dicki",
		MiddleName: "Darmawan",
		LastName:   "Saputra",
	}

	encoder.Encode(customer)
}
