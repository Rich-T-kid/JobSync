package middleware

import (
    "fmt"
    "net/http"
    "time"
)

func LogRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the current date and time
        now := time.Now()

        // Log the HTTP method, URL path, and current date and time
        fmt.Printf("[%s] %s at %s\n", now.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}

