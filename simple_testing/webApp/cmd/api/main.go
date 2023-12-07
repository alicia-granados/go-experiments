package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp/pkg/repository"
	"webapp/pkg/repository/dbrepo"

	"github.com/joho/godotenv"
)

const port = 8090

type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {

	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {

		log.Fatalf("Error cargando el archivo .env %v", err)
	}

	// Clave secreta para firmar el token (cámbiala por tu propia clave secreta)
	// Acceder a las variables de entorno
	secretKey := os.Getenv("SECRET_KEY")

	// Verificar si la variable está presente
	if secretKey == "" {
		log.Fatal("La variable SECRET_KEY no está configurada en el archivo .env")
	}

	var app application
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for application, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", secretKey, "signing secret")
	flag.Parse()

	conn, err := app.connectToDB()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Printf("starting api on port %d \n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
