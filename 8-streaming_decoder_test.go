package belajargojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
