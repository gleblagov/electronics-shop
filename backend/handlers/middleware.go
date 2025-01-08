package handlers

import (
	"net/http"

	"github.com/gleblagov/electronics-shop/utils"
)

func RoleMiddleware(role string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := utils.ValidateJwt(cookie.Value)
		if err != nil || claims.Role != role {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		handler(w, r)
	}
}
