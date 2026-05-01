package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	v := validator.New()
	type T struct {
		Zip string `validate:"postcode_iso3166_alpha2_field=Country"`
		Country string
	}
	err := v.Struct(T{Zip: "abc", Country: "BR"})
	fmt.Printf("Error: %v\n", err)
}
