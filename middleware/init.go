package middleware

import (
	"fmt"
	"net/http"
	"time"

	middleware "github.com/okhrimko/simplemiddleware/core"
)

var (
	// Default middleware
	Default = middleware.New(performanceLogger, addCORS, authHandler)
)

func performanceLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time1 := time.Now()
		next.ServeHTTP(w, r)
		time2 := time.Now()
		fmt.Printf("[%s] %q  %v\n", r.Method, r.URL.String(), time2.Sub(time1))
	}
}

func addCORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("addCORS: Adding CORS headers to response\n")
		next.ServeHTTP(w, r)
	}
}

func authHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("authHandler: Checking user\n")
		next.ServeHTTP(w, r)
	}
}
