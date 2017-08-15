package middleware

import (
	"net/http"
)

// Middleware struct
type Middleware struct {
	//ctx     context.Context
	chains  []middlewareHandler
	handler http.HandlerFunc
}

// MiddlewareHandler type
type middlewareHandler func(next http.Handler) http.HandlerFunc

// New created new middleware
func New(fns ...middlewareHandler) *Middleware {
	return &Middleware{
		//ctx:    context.Background(),
		chains: fns,
	}
}

// Then add main request handler
func (mw *Middleware) Then(fn func(w http.ResponseWriter, r *http.Request)) http.Handler {
	mw.handler = http.HandlerFunc(fn)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chainFn := mw.handler
		for _, itemFn := range mw.chains {
			chainFn = itemFn(chainFn)
		}
		chainFn(w, r)
	})
}
