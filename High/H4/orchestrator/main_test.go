package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNextTaskHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/next-task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код статуса: получено %v, ожидалось %v", status, http.StatusOK)
	}
}
