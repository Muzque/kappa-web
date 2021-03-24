package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
 }
 
func TestPingRoute(t *testing.T) {
	body := "{\"data\":{\"message\":\"pong\"}}"

	router := setupRouter()
	w := performRequest(router, "GET", "/ping")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, body, w.Body.String())
}
