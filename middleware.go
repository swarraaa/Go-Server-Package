package goServer

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Processing request %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func ConcurrencyMiddleware(maxConcurrency int) Middleware {
    sem := make(chan struct{}, maxConcurrency)

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            sem <- struct{}{} 
            defer func() { <-sem }()
            
            next.ServeHTTP(w, r)
        })
    }
}
