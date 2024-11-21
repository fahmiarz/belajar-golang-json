package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Fahmi",
		MiddleName: "Seazz",
		LastName:   "Arzalega",
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)
}
