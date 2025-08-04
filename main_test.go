package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {

	r:= setupRouter()

	w:= httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, Gin World!", w.Body.String())

}