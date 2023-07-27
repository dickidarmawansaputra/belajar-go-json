package belajargojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	// untuk melakukan encode gunakan json.Marshall()
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	logJson("Dicki")
	logJson(1)
	logJson(true)
	logJson([]string{"Dicki", "Darmawan", "Saputra"})
}
