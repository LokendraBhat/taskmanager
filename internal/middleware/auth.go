package middleware

import (
	"net/http"

	"crud-api/internal/config"
)

func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("user")
	if err != nil {
		return 0, err
	}
	var id int
	err = config.DB.QueryRow("SELECT id FROM users WHERE username = $1", cookie.Value).Scan(&id)
	return id, err
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUserID(r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}