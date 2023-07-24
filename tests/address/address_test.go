package address_test

import (
	. "testes/address"
	"testing"
)

type TestScenario struct {
	input    string
	expected string
}

func TestAddressType(t *testing.T) {
	// You can add t.Parallel() to run tests in parallel.
	// This is useful when you have a lot of tests.

	scenarios := []TestScenario{
		{"Rua das Flores", "Rua"},
		{"Avenida Cordilheiros dos Andes", "Avenida"},
		{"Praça das Rosas", "Praça"},
		{"Rua das Flores", "Rua"},
		{"Avenida dos Coqueiros", "Avenida"},
		{"Rua Lagoa da Prata", "Rua"},
		{"Estrada do Sol", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Avenida Automóvel Clube", "Avenida"},
		{"RUA ADELINO DE SOUZA CASTRO", "Rua"},
		{"AVENIDA DAS AMÉRICAS", "Avenida"},
		{"ESTRADA DO GALEÃO", "Estrada"},
		{"RODOVIA PRESIDENTE DUTRA", "Rodovia"},
		{"PRAÇA DA BANDEIRA", "Praça"},
		{"Prâça da Bandeira", "Invalid"},
		{"", "Invalid"},
		{"Central do CBLOL", "Invalid"},
	}

	for _, scenario := range scenarios {
		if received := AddressType(scenario.input); received != scenario.expected {
			t.Errorf("Expected %s, but received %s", scenario.expected, received)
		}
	}
}

func TestIncludes(t *testing.T) {
	fruits := []string{"apple", "banana", "orange"}

	input := "apple"
	expected := true

	if received := Includes(fruits, input); received != expected {
		t.Errorf("Expected %t, but received %t", expected, received)
	}
}
