package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	res := greeting("Code.education Rocks!")
	if res != "<b>Code.education Rocks!</b>" {
		t.Errorf("String formatada incorretamente: %s, esperava: %s.", res, "<b>Code.education Rocks!</b>")
	}

	// http request test
	t.Run("http request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		handler(response, request)

		got := response.Body.String()
		want := "<b>Code.education Rocks!</b>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
