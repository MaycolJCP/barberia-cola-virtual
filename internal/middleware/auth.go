package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Clave secreta utilizada para firmar y validar los tokens JWT.
// Debe ser la misma que se usa al momento de generar el token durante el login.
var jwtSecret = []byte("clave-secreta-barberia")

// Se define un tipo personalizado para usar como clave dentro del contexto.
// Esto evita conflictos con otras claves de tipo string.
type contextKey string

// Claves utilizadas para guardar información del usuario en el contexto.
const (
	UsuarioIDKey contextKey = "usuario_id"
	CorreoKey    contextKey = "correo"
	RolKey       contextKey = "rol"
)

// Aqui valida que la petición tenga un token JWT válido.
// Si el token es correcto, guarda los datos del usuario en el contexto
// para que los handlers puedan utilizarlos posteriormente.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Obtiene el encabezado Authorization enviado por el cliente.
		authHeader := r.Header.Get("Authorization")

		// Verifica que el encabezado exista.
		if authHeader == "" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}

		// Verifica que el formato sea "Bearer <token>".
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Formato de token invalido", http.StatusUnauthorized)
			return
		}

		// Elimina la palabra "Bearer " para quedarse únicamente con el JWT.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Valida el token utilizando la clave secreta.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Si el token es inválido o ocurrió un error, se rechaza la petición.
		if err != nil || !token.Valid {
			http.Error(w, "Token invalido", http.StatusUnauthorized)
			return
		}

		// Convierte los datos del token (Claims) a un mapa para poder acceder a ellos.
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims invalidos", http.StatusUnauthorized)
			return
		}

		// Guarda la información del usuario dentro del contexto de la petición.
		// Así cualquier handler podrá obtener estos datos sin volver a leer el token.
		ctx := context.WithValue(r.Context(), UsuarioIDKey, claims["usuario_id"])
		ctx = context.WithValue(ctx, CorreoKey, claims["correo"])
		ctx = context.WithValue(ctx, RolKey, claims["rol"])

		// Continúa hacia el siguiente middleware o handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminOnly verifica que el usuario autenticado tenga el rol ADMIN.
// Si no lo tiene, devuelve un error 403 (Forbidden).
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Obtiene el rol almacenado previamente por AuthMiddleware.
		rol := r.Context().Value(RolKey)

		// Solo los administradores pueden continuar.
		if rol != "ADMIN" {
			http.Error(w, "Acceso denegado: solo ADMIN", http.StatusForbidden)
			return
		}

		// Continúa con el siguiente handler.
		next.ServeHTTP(w, r)
	})
}

// Aqui verifica que el usuario autenticado tenga el rol CLIENTE.
// Si el rol no coincide, devuelve un error 403.
func ClienteOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Obtiene el rol guardado en el contexto.
		rol := r.Context().Value(RolKey)

		// Solo los clientes pueden acceder.
		if rol != "CLIENTE" {
			http.Error(w, "Acceso denegado: solo CLIENTE", http.StatusForbidden)
			return
		}

		// Continúa con el siguiente handler.
		next.ServeHTTP(w, r)
	})
}