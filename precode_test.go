package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("Get", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)

	status := responseRecorder.Code

	assert.Equal(t, http.StatusOK, status)
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("Get", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	body := responseRecorder.Body

	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, body)
}

func TestСityThatDoesntExist(t *testing.T) {

	req := httptest.NewRequest("Get", "/cafe?count=4&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, http.StatusBadRequest, status)

	body := responseRecorder.Body.String()

	assert.Equal(t, body, "wrong city value")
}
