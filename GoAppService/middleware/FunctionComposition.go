package middleware

import "net/http"

func ValidateTokenAdminAccess(next http.HandlerFunc) http.HandlerFunc {
	return checkToke(checkAdminAccess(checkLoginStatus(next)))
}

func ValidateTokenUserAccess(next http.HandlerFunc) http.HandlerFunc {
	return checkToke(checkUserAccess(checkLoginStatus(next)))
}
