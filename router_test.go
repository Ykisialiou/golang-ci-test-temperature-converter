package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
        "encoding/json"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Ok", w.Body.String())
}

func TestConvertTemperature(t *testing.T) {
       router := setup()
       //var jsonData temperature
       //router.POST("/temperature", convertTemperature)
       //jsonStr := "{'TemperatureC' : '30'}"
       //jsonData = {"TemperatureC", "30"}
       data := &temperature{TemperatureC: "30"}
       jsonValue, _ := json.Marshal(data)

       req, _ := http.NewRequest("POST", "/temperature", bytes.NewBuffer(jsonValue))
       w := httptest.NewRecorder()
       router.ServeHTTP(w, req)
       assert.Equal(t, "86", w.Body.String())

}

