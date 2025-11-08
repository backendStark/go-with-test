package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

func TestHandler_ReturnsPlainText(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code equal 200, but got: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Errorf("Expected header 'Content-Type' equal 'text/plain; charset=utf-8', but got: %q", resp.Header.Get("Content-Type"))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	body := strings.TrimSpace(string(bodyBytes))

	ts, err := time.Parse(time.RFC3339, body)

	if err != nil {
		t.Fatalf("Failed to parse RFC3339 time from body %q: %v", body, err)
	}

	if ts.Weekday().String() == "" {
		t.Error("Expected non-empty weekday")
	}
	if ts.Day() == 0 {
		t.Error("Expected non-zero day of month")
	}
	if ts.Month().String() == "" {
		t.Error("Expected non-empty month")
	}
	if ts.Year() == 0 {
		t.Error("Expected non-zero year")
	}
	if ts.Hour() == 0 && ts.Minute() == 0 && ts.Second() == 0 {
		t.Error("Expected non-zero time components (hour/minute/second)")
	}
}

func TestHandler_RejectsNonGETMethods(t *testing.T) {
	testCases := []struct {
		name   string
		method string
	}{
		{name: "POST method", method: "POST"},
		{name: "PUT method", method: "PUT"},
		{name: "DELETE method", method: "DELETE"},
		{name: "PATCH method", method: "PATCH"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "http://localhost:8080", nil)
			w := httptest.NewRecorder()

			t.Logf("Testing method: %s", req.Method)

			handler(w, req)

			resp := w.Result()
			if resp.StatusCode != 405 {
				t.Errorf("Expected status 405 for %s, but got: %d", tc.method, resp.StatusCode)
			}
		})
	}
}

func TestClientIP_ExtractFromXForwardedFor(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.1, 198.51.100.1, 192.0.2.1")
	ip := clientIP(req)
	expected := "203.0.113.1"
	if ip != expected {
		t.Errorf("Expected IP %q, but got: %q", expected, ip)
	}
}
