package main

import "fmt"

func main() {

	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
		"Alicia": 25,
		"Juana":  10,
	}

	fmt.Println(m)

	fmt.Println(m["Batman"])
	fmt.Println(m["Eduardo"])

	v, ok := m["Eduardo"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Eduardo"]; ok {
		fmt.Println("Imprimiendo desde el if", v)
	}

	if v, ok := m["Robin"]; ok {
		fmt.Println("Imprimiendo desde el if 2", v)
	}

	/* agregar elemento  y Range*/
	fmt.Println("")
	m["Eduardo"] = 31

	for llave, valor := range m {
		fmt.Println(llave, valor)
	}

	/*borrar elementos de un mapa*/
	fmt.Println("")

	delete(m, "Robin")
	fmt.Println(m)

	if v, ok := m["Robin"]; ok {
		fmt.Println("Se borró la llave con el valor: ", v)
		delete(m, "Robin")
	}
	fmt.Println(m)

	if v, ok := m["Juana"]; ok {
		fmt.Println("Se borró la llave con el valor: ", v)
		delete(m, "Juana")
	}

	fmt.Println(m)
}
