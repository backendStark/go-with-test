package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestHandler_ReturnsJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code equal 200, but got: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("Expected header 'Content-Type' equal 'application/json; charset=utf-8', but got: %q", resp.Header.Get("Content-Type"))
	}

	var result timeJSON

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response json: %v", err)
	}

	if result.DayOfWeek == "" {
		t.Error("Expected non-empty DayOfWeek")
	}
	if result.DayOfMonth == 0 {
		t.Error("Expected non-zero DayOfMonth")
	}
	if result.Month == "" {
		t.Error("Expected non-empty Month")
	}
	if result.Year == 0 {
		t.Error("Expected non-zero Year")
	}
	if result.Hour == 0 && result.Minute == 0 && result.Second == 0 {
		t.Error("Expected non-zero time fields (Hour, Minute, Second)")
	}
}
