package belajargojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"Dicki","MiddleName":"Darmawan","LastName":"Saputra"}`
	jsonBytes := []byte(jsonString)

	// jika tidak pake pointer, datanya diubah di unmarshal. nantinya data variable ini tidak mendapat peruahannya
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
