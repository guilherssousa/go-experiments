package address

import "testing"

func TestAddressType(t *testing.T) {
	address := "Avenida Cordilheiros dos Andes"

	expected := "Avenida"

	if received := AddressType(address); received != expected {
		t.Errorf("Expected %s, but received %s", expected, received)
	}
}

func TestIncludes(t *testing.T) {
	fruits := []string{"apple", "banana", "orange"}

	fruit := "apple"

	expected := true

	if received := Includes(fruits, fruit); received != expected {
		t.Errorf("Expected %t, but received %t", expected, received)
	}
}
