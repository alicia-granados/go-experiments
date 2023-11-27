package main

import (
	"os"
	"testing"
)

//create an app var of type application
var app application

func TestMain(m *testing.M) {

	os.Exit(m.Run())

}
