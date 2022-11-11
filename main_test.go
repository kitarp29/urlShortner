package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestRedirect(t *testing.T) {
	// Setup
	for key, value := range ListMap {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?u="+key, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Assertions
		got := Redirect(c)

		if got != value {
			t.Errorf("Redirect() = %q, want %q", got, value)
		} else {
			t.Logf("FAILED!")
		}
	}
}

func TestAddLink(t *testing.T) {
	// Setup

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/addlink?link=https://www.twitter.com", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	got := AddLink(c)
	chk := false

	for key := range ListMap {
		if got == "http://localhost:8000/?u="+key {
			chk = true
			break
		}
	}
	if !chk {
		t.Errorf("AddLink() = %q, was not found", got)
	}

}
