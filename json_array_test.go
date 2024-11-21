package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Fahmi",
		MiddleName: "Seazz",
		LastName:   "Arzalega",
		Age:        25,
		Gender:     "Laki-Laki",
		Married:    false,
		Hobbies:    []string{"Coding", "Reading", "Gaming"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {

	jsonString := `{"FirstName":"Fahmi","MiddleName":"Seazz","LastName":"Arzalega","Age":25,"Gender":"Laki-Laki","Married":false,"Hobbies":["Coding","Reading","Gaming"]}`
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
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Fahmi",
		Addresses: []Address{
			{
				Street:     "Jl.Buntu",
				Country:    "Indonesia",
				PostalCode: "11530",
			},
			{
				Street:     "Jl.Buntu lagi",
				Country:    "Indonesia",
				PostalCode: "9090",
			},
		},
		Hobbies: []string{"Coding", "Reading", "Gaming"},
		Age:     20,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"FirstName":"Fahmi","MiddleName":"","LastName":"","Age":20,"Gender":"","Married":false,"Hobbies":["Coding","Reading","Gaming"],"Addresses":[{"Street":"Jl.Buntu","Country":"Indonesia","PostalCode":"11530"},{"Street":"Jl.Buntu lagi","Country":"Indonesia","PostalCode":"9090"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {

	jsonString := `[{"Street":"Jl.Buntu","Country":"Indonesia","PostalCode":"11530"},{"Street":"Jl.Buntu lagi","Country":"Indonesia","PostalCode":"9090"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)

	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{{
		Street:     "Jl.Buntu",
		Country:    "Indonesia",
		PostalCode: "11530",
	},
		{
			Street:     "Jl.Buntu lagi",
			Country:    "Indonesia",
			PostalCode: "9090",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
