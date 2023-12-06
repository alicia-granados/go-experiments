package main

//maneja la obtención y verificación de tokens de autenticación JWT (JSON Web Token) a partir de encabezados de solicitud HTTP.
import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//Define dos constantes que representan el tiempo de expiración de los tokens JWT y de renovación,
const jwtTokenExpiry = time.Minute * 15
const refreshTokenExpiry = time.Hour * 24

// para contener un par de tokens (un token de acceso y un token de actualización)
type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//representar los reclamos incluidos en los tokens JWT
type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

// se encarga de manejar la extracción y verificación del token de acceso de las solicitudes HTTP.
func (app *application) getTokenFromHeaderandVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {

	//we expect our authorization header to look like this:
	//Bearer <token>

	// add a header
	w.Header().Add("Vary", "Authorization")

	//get the authorization header
	authHeader := r.Header.Get("Authorization")

	//sanity check
	if authHeader == "" {
		return "", nil, errors.New("no auth header")
	}

	//split the header on spaces
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid auth header")
	}

	//check to see if we have the word "bearer"
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("unauthorized: no bearer")
	}

	token := headerParts[1]

	//declare an empty claims variable
	claims := &Claims{}

	//parse the token with our claims( we read into claims), using our secret (from the receiver)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexecpected signing method: %v", token.Header["alg"])
		}

		return []byte(app.JWTSecret), nil
	})

	//check for an error; note that this catches expired tokens as well
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("expired token")
		}
		return "", nil, err
	}

	//make sure that we issued this token
	if claims.Issuer != app.Domain {
		{
			return "", nil, errors.New("incorrect issuer")
		}
	}

	//valid token
	return token, claims, nil
}
