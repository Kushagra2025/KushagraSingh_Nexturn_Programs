// middleware/error_handler.go
package middleware

import (
	"encoding/json"
	"net/http"
)

// ErrorHandler middleware catches errors and formats them as JSON responses
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use a deferred function to catch panics and return an error as JSON
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Internal Server Error",
				})
			}
		}()
		next.ServeHTTP(w, r) // Call the next handler
	})
}
