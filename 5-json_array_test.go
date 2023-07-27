package belajargojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Dicki",
		MiddleName: "Darmawan",
		LastName:   "Saputra",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Dicki","MiddleName":"Darmawan","LastName":"Saputra","Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Dicki",
		MiddleName: "Darmawan",
		LastName:   "Saputra",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan",
				Country:    "ID",
				PostalCode: "00",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Dicki","MiddleName":"Darmawan","LastName":"Saputra","Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jalan","Country":"ID","PostalCode":"00"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan",
			Country:    "ID",
			PostalCode: "00",
		},
		{
			Street:     "Jalan 2",
			Country:    "ID 2",
			PostalCode: "00",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan","Country":"ID","PostalCode":"00"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
