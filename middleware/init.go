package middleware

import (
	"fmt"
	"net/http"
	"time"

	middleware "github.com/okhrimko/simplemiddleware/core"
)

var (
	// General middleware
	Default = middleware.New(performanceLogger, addCORS, authHandler)
)

func performanceLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time1 := time.Now()
		next.ServeHTTP(w, r)
		time2 := time.Now()
		fmt.Fprintf(w, "[%s] %q  %v<br/>", r.Method, r.URL.String(), time2.Sub(time1))
	}
}

func addCORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "addCORS: Adding CORS headers to response<br/>")
		next.ServeHTTP(w, r)
	}
}

func authHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "authHandler: Checking user<br/>")
		next.ServeHTTP(w, r)
	}
}
