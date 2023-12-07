package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBOdy        string
		expectedStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com", "password":"secret"}`, http.StatusOK},
		{"not json", `Im not json`, http.StatusUnauthorized},
		{"empty json ", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":""}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com"}`, http.StatusUnauthorized},
		{"invalid user", `{"email":"admin@someotherdomain.com", "password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		var reader = strings.NewReader(e.requestBOdy)
		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)

		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code; expected %d but gor %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}
