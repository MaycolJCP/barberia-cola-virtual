package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("clave-secreta-barberia")

type contextKey string

const (
	UsuarioIDKey contextKey = "usuario_id"
	CorreoKey    contextKey = "correo"
	RolKey       contextKey = "rol"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Formato de token invalido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token invalido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims invalidos", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UsuarioIDKey, claims["usuario_id"])
		ctx = context.WithValue(ctx, CorreoKey, claims["correo"])
		ctx = context.WithValue(ctx, RolKey, claims["rol"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rol := r.Context().Value(RolKey)

		if rol != "ADMIN" {
			http.Error(w, "Acceso denegado: solo ADMIN", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ClienteOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rol := r.Context().Value(RolKey)

		if rol != "CLIENTE" {
			http.Error(w, "Acceso denegado: solo CLIENTE", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}