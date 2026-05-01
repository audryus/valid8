package valid8_test

import (
	"testing"

	"github.com/audryus/valid8"
)

type User struct {
	Name    string  `validate:"required"`
	Email   string  `validate:"required,email"`
	Age     int     `validate:"required,gte=18"`
	Address Address `validate:"required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required,alphanumspace"`
}

func TestValid8(t *testing.T) {
	v := valid8.New(valid8.WithLocales(valid8.PT_BR))
	if v == nil {
		t.Error("New() returned nil")
	}

	u := User{
		Address: Address{
			City: "%",
		},
	}
	err := v.Struct(u, valid8.PT_BR)

	if len(err) == 0 {
		t.Error("No errors are impossible here")
	}

	errs := valid8.ErrorsToMap(err)

	containsEquals(t, errs, "user.name", "Name é um campo obrigatório")
	containsEquals(t, errs, "user.email", "Email é um campo obrigatório")
	containsEquals(t, errs, "user.age", "Age é um campo obrigatório")
	containsEquals(t, errs, "user.address.street", "Street é um campo obrigatório")
	containsEquals(t, errs, "user.address.city", "City deve conter apenas caracteres alfanuméricos e espaços")
}

func containsEquals(t *testing.T, mep map[string]string, key, value string) {
	v, ok := mep[key]
	if !ok {
		t.Errorf("Key %s not found", key)
	} else if v != value {
		t.Errorf("Founded %s, expected %s", v, value)
	}
}
