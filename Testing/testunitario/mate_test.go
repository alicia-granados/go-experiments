package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("SUma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}
