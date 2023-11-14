package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // PIEDRA. VENCE A TIJERAS . (TIJERAS +1) %3 =0
	PAPER    = 1 // PAPEL. VENCE A PIEDRA (PIEDRA +1) % 3 = 1
	SCISSORS = 2 // TIJERAS. VENCEN AL PAPEL (PAPEL +1) %3 = 2
)

// ESTRUCTURA PARA DAR RESULTADO DE CADA RONDA
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// mensjaes para cuando gana
var winMessages = []string{
	"¡Bien hecho!",
	"¡Buen trabajo!",
	"Deberias comprar un boleti¿o de loteria",
}

// mensjaes para cuando pierde
var loseMessages = []string{
	"¡Que lastima!",
	"¡Intentalo de nuevo!",
	"Hoy simplemente no es tu dia",
}

// mesajes de empate
var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"oh oh.Intentalo de nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

// variables para el puntuaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	//mensaje dependiendo de lo que eligio la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio piedra"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"
	}

	// Generar un numero aleatorio de 0-2 , que usamos para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)
	//declarar una var para contener el mensaje
	var message string

	if playerValue == computerValue {

		roundResult = "Es un empate"

		// seleccionar mensjae de drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "¡El jugador gana!"
		//seleccionar mensaje de winMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		//seleccionar mensaje de loseMessages
		message = loseMessages[messageInt]
	}

	// retornar el resultado
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
