package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(io.Reader(reader))

	customer := &Customer{}
	decoder.Decode(customer)
	
	fmt.Println(customer)
}
