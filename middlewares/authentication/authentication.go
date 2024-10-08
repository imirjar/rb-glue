package authentication

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/imirjar/rb-glue/models"
)

// AuthServiceResponse представляет структуру ответа от сервиса авторизации
type AuthServiceResponse struct {
	Valid  bool     `json:"valid"`
	User   string   `json:"user,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Roles  []string `json:"roles,omitempty"`

	// Добавьте другие поля по необходимости
}

func Authenticate(authServiceURL string) models.Middleware {
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

			token = strings.TrimPrefix(token, "Bearer ")

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

			// Читаем и декодируем ответ от сервиса авторизации
			var authResp AuthServiceResponse
			if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if !authResp.Valid {
				http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
				return
			}

			// Можно добавить информацию о пользователе в контекст, если необходимо
			// ctx := context.WithValue(r.Context(), "user", authResp.User)
			// next.ServeHTTP(w, r.WithContext(ctx))

			// Переходим к следующему обработчику
			next.ServeHTTP(w, r)
		})
	}
}
