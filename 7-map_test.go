package belajargojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"001","name":"Product 1","image_url":"image file"}`
	jsonBytes := []byte(jsonString)

	// gunakan map jika jsonnya dinamis (atribut berubah-rubah), kalo pake struct akan sulit
	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["name"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "001",
		"name":  "Tes",
		"price": 1000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
