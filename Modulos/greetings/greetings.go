package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//hello devuelve un saludo para la persona especifica

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("NOmbre vacio")
	}
	//devuelve un slaudo que incluye el nombre del mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// hellos returns a map that each of the named peole with a greeting message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello((name))
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil

}

// randomfromat devuelve uno de varios formatos de mensajes
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"Hola , %v, bienvenudo",
		"¡que bueno verte %v!",
		"saludo %v, encantado de conocerte",
	}

	return formats[rand.Intn(len(formats))]
}
