package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("SUma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}

//nueva forma de testeo

func TestSuma2(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("SUma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

/*coverage : porcentaje del testeo go test-cover
o crear archivo  go test -coverprofile=coverage.out ,
 ver que funciones estan cubiertas go tool cover -func=coverage.out
 go tool cover -html=coverage.out muestra en el html las funciones que ya han sido cubiertas de color verde

PROFILING: saver donde se esta demorando
 go test -cpuprofile=cou.out
 go tool pprof cou.out
 dentrro del archivo poner :
 top
 list Fibonacci
 web o pdf
 quit salir

 obtemer reportes install :  sudo apt install graphviz

*/

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{25, 250, 250},
	}

	for _, item := range tabla {
		num_may := GetMax(item.a, item.b)

		if num_may != item.c {
			t.Errorf("Sel numero mayor , es %d se esperaba %d", num_may, item.c)
		}
	}
}

func TestFibo(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fib := Fibonacci(item.a)

		if fib != item.b {
			t.Errorf("Fibonacci , es %d se esperaba %d", fib, item.b)
		}
	}
}
