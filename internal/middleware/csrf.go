package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const csrfCookieName = "csrf_token"
const CSRFFormField = "csrf_token"

func generateCSRFToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// GetOrSetCSRFToken returns the existing CSRF token from the cookie, or creates
// and sets a new one. Call this when rendering any page that has a form.
func GetOrSetCSRFToken(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(csrfCookieName)
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}
	token := generateCSRFToken()
	http.SetCookie(w, &http.Cookie{
		Name:     csrfCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return token
}

// CSRFMiddleware validates the CSRF token on every POST request using the
// double-submit cookie pattern: the form field must match the cookie value.
func CSRFMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			cookie, err := r.Cookie(csrfCookieName)
			if err != nil || cookie.Value == "" {
				http.Error(w, "Forbidden: missing CSRF cookie", http.StatusForbidden)
				return
			}
			if r.FormValue(CSRFFormField) != cookie.Value {
				http.Error(w, "Forbidden: invalid CSRF token", http.StatusForbidden)
				return
			}
		}
		next(w, r)
	}
}
