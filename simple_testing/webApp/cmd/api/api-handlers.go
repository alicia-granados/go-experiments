package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	//read a json payload
	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//look up the user by email address
	user, err := app.DB.GetUserByEmail(creds.Username)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//generate tokens
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//send token to user
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)

}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

	// ParseForm analiza la URL de la solicitud, el cuerpo o los campos de la forma
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Obtener el token de actualización del formulario de la solicitud
	refreshToken := r.Form.Get("refresh_token")

	// Crear una estructura para almacenar las reclamaciones del token
	claims := &Claims{}

	// Parsear el token de actualización utilizando jwt.ParseWithClaims
	_, err = jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})

	if err != nil {
		// Si hay un error al analizar el token, devolver un error JSON
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Si hay un error al analizar el token, devolver un error JSON
	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		app.errorJSON(w, errors.New("refresh token does not need renewed yet"), http.StatusTooEarly)
	}

	// Obtener el ID de usuario de las reclamaciones
	userID, err := strconv.Atoi(claims.Subject)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Obtener información del usuario desde la base de datos utilizando el ID
	user, err := app.DB.GetUser(userID)

	if err != nil {
		app.errorJSON(w, errors.New("unknown user"), http.StatusBadRequest)
		return
	}

	// Generar un nuevo par de tokens de acceso y actualización
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Establecer una cookie HTTP para el nuevo token de actualización
	http.SetCookie(w, &http.Cookie{
		Name:     "__Host-refrresh-token",
		Path:     "/",
		Value:    tokenPairs.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	})

	// Escribir la respuesta JSON con el nuevo par de tokens
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {

}
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

}
func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {

}
