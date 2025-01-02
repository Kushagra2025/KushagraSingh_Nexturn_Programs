package middleware

import (
	"encoding/json"
	"net/http"
)

func ValidateProductInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Basic validation
		if product["name"] == nil || product["price"] == nil {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
