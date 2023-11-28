package main

import (
	"os"
	"testing"
)

//create an app var of type application
var app application

func TestMain(m *testing.M) {

	app.Session = getSession()

	os.Exit(m.Run())

}
