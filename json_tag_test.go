package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac Book Pro",
		Price:    25000000,
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac Book Pro","price":25000000,"image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
