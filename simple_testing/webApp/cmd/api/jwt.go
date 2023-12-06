package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func generarJWT() string {

	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {

		log.Fatalf("Error cargando el archivo .env", err)
	}

	// Clave secreta para firmar el token (cámbiala por tu propia clave secreta)
	// Acceder a las variables de entorno
	secretKey := os.Getenv("SECRET_KEY")

	// Imprimir contenido para verificar
	//fmt.Println("SECRET_KEY:", secretKey)

	// Verificar si la variable está presente
	if secretKey == "" {
		log.Fatal("La variable SECRET_KEY no está configurada en el archivo .env")
	}

	// Crear el encabezado (header)
	header := jwt.MapClaims{
		"alg": "HS256", // Algoritmo de firma (puedes cambiarlo según tus necesidades)
		"typ": "JWT",
	}

	// Crear la carga útil (payload)
	payload := jwt.MapClaims{
		"sub":  "1234567890", // Sujeto (identificación del usuario)
		"name": "John Doe",
		"iat":  time.Now().Unix(),                       // Tiempo de emisión del token
		"exp":  time.Now().Add(time.Minute * 15).Unix(), // Tiempo de expiración del token (30 minutos en este caso)
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"header":  header,
		"payload": payload,
	})

	// Firma del token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error al firmar el token:", err)
		return "Error al firmar el token:"
	}

	// Imprimir el token generado
	//fmt.Println("Token JWT:", tokenString)
	return tokenString
}
