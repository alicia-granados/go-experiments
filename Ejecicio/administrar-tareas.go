package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}
type ListaTareas struct {
	tareas []Tarea
}

// metodo para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para marcar como completado
func (l *ListaTareas) marcarCompletadp(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tarea
func (l *ListaTareas) editarTareas(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para elimianr tarea
func (l *ListaTareas) eliminarTareas(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//isntancia de lista de tareas
	lista := ListaTareas{}

	//instancia de bufio para entrad de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("seleccione una opcion\n",
			"1.Agregar  tarea \n",
			"2.Marcar tarea como completada\n",
			"3.editar tarea\n",
			"4.Eliminar tarea\n",
			"5.Salir")

		fmt.Println("Ingrese la opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln((&index))
			lista.marcarCompletadp((index))
			fmt.Println("tarea marcada como compeltada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Println("ingrese el indice de la tarea qie desea actualizar ")
			fmt.Scanln(&index)
			fmt.Println("ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTareas(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea eliminar")
			fmt.Scanln(&index)
			lista.eliminarTareas(index)
			fmt.Println("tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opcion invalida")
		}

		//LISTAR TODAS LAS TAREAS
		fmt.Println("lista de tareas")
		fmt.Println("-----------------------------")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - COmpletado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("----------------------------")
	}
}
