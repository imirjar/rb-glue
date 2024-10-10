package authentication

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Authenticate(authServiceURL string, user User) func(http.Handler) http.Handler {
	client := &http.Client{
		Timeout: 5 * time.Second, // Устанавливаем таймаут для запросов к сервису авторизации
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized: no token provided", http.StatusUnauthorized)
				return
			}

			// token = strings.TrimPrefix(token, "Bearer ")

			// Создаем запрос к сервису авторизации
			req, err := http.NewRequest("POST", authServiceURL, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Добавляем токен в заголовки запроса
			req.Header.Set("Authorization", token)

			// Отправляем запрос к сервису авторизации
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

			user := User{}
			if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if user.ID != "" {
				if user.ID != user.ID {
					http.Error(w, "Unauthorized: Wrong user", http.StatusForbidden)
					return
				}
			}

			if len(user.Groups) >= 0 {
				for _, g := range user.Groups {
					if !user.hasGroup(g) {
						http.Error(w, "Unauthorized: Roles or Groups are missing", http.StatusForbidden)
						return
					}
				}
			}

			if len(user.Roles) >= 0 {
				for _, r := range user.Roles {
					if !user.hasRole(r) {
						http.Error(w, "Unauthorized: Roles or Groups are missing", http.StatusForbidden)
						return
					}
				}

			}

			next.ServeHTTP(w, r)
		})
	}
}
