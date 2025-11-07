package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type accessLog struct {
	IP        string    `json:"ip"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	UserAgent string    `json:"user_agent"`
	Time      time.Time `json:"time"`
}

type timeJSON struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)

		entry := accessLog{
			IP:        ip,
			Method:    r.Method,
			Path:      r.URL.Path,
			UserAgent: r.UserAgent(),
			Time:      time.Now(),
		}

		b, err := json.Marshal(entry)
		if err != nil {
			log.Printf(`{"level":"error","msg":"marshal log failed","err":%q}`, err)
		} else {
			log.Println(string(b))
		}

		next.ServeHTTP(w, r)
	})
}

func clientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		ip := strings.TrimSpace(parts[0])
		if ip != "" {
			return ip
		}
	}

	if xrip := strings.TrimSpace(r.Header.Get("X-Real-IP")); xrip != "" {
		return xrip
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return host
	}

	return r.RemoteAddr
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	timeNow := time.Now()

	accept := r.Header.Get("Accept")

	if strings.Contains(accept, "application/json") {
		payload := timeJSON{
			DayOfWeek:  timeNow.Weekday().String(),
			DayOfMonth: timeNow.Day(),
			Month:      timeNow.Month().String(),
			Year:       timeNow.Year(),
			Hour:       timeNow.Hour(),
			Minute:     timeNow.Minute(),
			Second:     timeNow.Second(),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, "failed to encode json", http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte(timeNow.Format(time.RFC3339)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	logged := loggingMiddleware(mux)
	log.Println(`{"level":"info","msg":"listening","addr":":8080"}`)
	log.Fatal(http.ListenAndServe(":8080", logged))
}
