package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuth is a simple Basic Authentication middleware
func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the "Authorization" header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// The "Authorization" header should be in the format "Basic <credentials>"
		// Example: Basic dXNlcjpwYXNzd29yZA==
		if !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Extract the base64-encoded part after "Basic "
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			http.Error(w, "Invalid base64 encoding", http.StatusUnauthorized)
			return
		}

		// The decoded credentials are in the format "username:password"
		credentials := strings.SplitN(string(decodedCredentials), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid Authorization credentials", http.StatusUnauthorized)
			return
		}

		username, password := credentials[0], credentials[1]

		// Check if the username and password are correct
		if username != "admin" || password != "password" { // Replace these values with your valid credentials
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// If authentication passes, call the next handler
		next.ServeHTTP(w, r)
	})
}
