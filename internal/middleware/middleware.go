package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO validate user
		// temp
		if true {
			fmt.Println("User is Authenticated")
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// TODO
// Rate limiter
// Logging Middleware
// Error middleware
// metrics
