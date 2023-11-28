package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func getSession() *scs.SessionManager {
	//crear una nueva sesión
	session := scs.New()
	//luego configurando algunas opciones como la duración de la sesión
	session.Lifetime = 24 * time.Hour
	// persistencia de la cookie
	session.Cookie.Persist = true
	// el modo SameSite de la cookie
	session.Cookie.SameSite = http.SameSiteLaxMode
	//la seguridad de la cookie
	session.Cookie.Secure = true
	return session
}
