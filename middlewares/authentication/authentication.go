package authentication

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Compare JWT token payload with expected user ID, Groups and Roles
// if no expected -> just validate the JWT token on the auth authService
func Authenticate(url string, params UserParams) func(http.Handler) http.Handler {

	// Custom client
	client := &http.Client{
		Timeout: 5 * time.Second, // Устанавливаем таймаут для запросов к сервису авторизации
	}

	// Return auth middleware
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized: no token provided", http.StatusUnauthorized)
				return
			}

			// Make request to auth service
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Add token in header
			req.Header.Set("Authorization", token)

			// Make req
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				http.Error(w, "Unauthorized: auth service error", http.StatusUnauthorized)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
				return
			}

			if err := json.NewDecoder(resp.Body).Decode(&params); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// If a user must be the user
			if params.ID != "" {
				if params.ID != params.ID {
					http.Error(w, "Unauthorized: Wrong user", http.StatusForbidden)
					return
				}
			}

			// If user must be at least in one of the groups
			if len(params.Groups) >= 0 {
				for _, g := range params.Groups {
					if !params.hasGroup(g) {
						http.Error(w, "Unauthorized: Roles or Groups are missing", http.StatusForbidden)
						return
					}
				}
			}

			// If user must has at least in one of the groups
			if len(params.Roles) >= 0 {
				for _, r := range params.Roles {
					if !params.hasRole(r) {
						http.Error(w, "Unauthorized: Roles or Groups are missing", http.StatusForbidden)
						return
					}
				}

			}

			next.ServeHTTP(w, r)
		})
	}
}
