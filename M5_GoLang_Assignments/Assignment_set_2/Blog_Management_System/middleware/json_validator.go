package middleware

import (
	"encoding/json"
	"net/http"
)

func ValidateJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Invalid COntent-type,must be application/json", http.StatusBadRequest)
			return
		}

		var js map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&js); err != nil {
			http.Error(w, "Ivalid JSON payload", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
