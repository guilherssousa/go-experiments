package shapes_test

import (
	"math"
	. "nice-testing/shapes"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10.0, 12.0}
		expected := float64(120)

		if received := ret.Area(); received != expected {
			t.Fatalf("expected '%.2f' but got '%.2f'", expected, received)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		expected := float64(math.Pi * 100)

		if received := circ.Area(); received != expected {
			t.Fatalf("expected '%.2f' but got '%.2f'", expected, received)
		}
	})
}
