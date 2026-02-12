package sum

import "testing"

func TestSum(t *testing.T) {
	// Preparare le variabile per il testing
	var num1, num2 int = 2, 3

	// Invocare la funzione
	res := sum(num1, num2)

	// Controllare il risultato
	if res != 5 {
		t.Errorf("Expected 5, got %d", res)
	}
}