package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Gender     string
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Fahmi",
		MiddleName: "Seazzz",
		LastName:   "Arzalega",
		Age:        25,
		Gender:     "Laki-Laki",
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Fahmi","MiddleName":"Seazzz","LastName":"Arzalega","Age":25,"Gender":"Laki-Laki","Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Gender)
	fmt.Println(customer.Married)
}
