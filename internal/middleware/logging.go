package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	var loggingFunction http.HandlerFunc
	t := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%v - %v - %v\n", r.URL.Path, r.Method, time.Now().Format("15:04:05 2006-01-02"))
		next.ServeHTTP(w, r)
	}
	loggingFunction = t
	return loggingFunction
}
