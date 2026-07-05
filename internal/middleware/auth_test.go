package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
)

// =======================================
// TEST MIDDLEWARE SIN TOKEN
//
// PREPARAR:
// Crear una petición sin Authorization.
//
// EJECUTAR:
// Pasarla por AuthMiddleware.
//
// VERIFICAR:
// Debe responder 401 Unauthorized.
// =======================================
func TestAuthMiddleware_SinToken(t *testing.T) {

	// Handler de prueba protegido por el middleware
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Request sin header Authorization
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Recorder captura la respuesta sin levantar servidor real
	rec := httptest.NewRecorder()

	// Ejecutar petición falsa
	handler.ServeHTTP(rec, req)

	// Verificar status esperado
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("se esperaba status %d pero llegó %d", http.StatusUnauthorized, rec.Code)
	}
}

// =======================================
// TEST MIDDLEWARE TOKEN INVALIDO
//
// PREPARAR:
// Crear una petición con token falso.
//
// EJECUTAR:
// Pasarla por AuthMiddleware.
//
// VERIFICAR:
// Debe responder 401 Unauthorized.
// =======================================
func TestAuthMiddleware_TokenInvalido(t *testing.T) {

	// Handler de prueba
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Request con token falso
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	req.Header.Set(
		"Authorization",
		"Bearer token-falso",
	)

	rec := httptest.NewRecorder()

	// Ejecutar
	handler.ServeHTTP(rec, req)

	// Verificar
	if rec.Code != http.StatusUnauthorized {
		t.Errorf(
			"se esperaba %d pero llegó %d",
			http.StatusUnauthorized,
			rec.Code,
		)
	}
}

// =======================================
// TEST ADMIN ONLY CON ROL CLIENTE
//
// PREPARAR:
// Crear una petición con rol CLIENTE.
//
// EJECUTAR:
// Pasarla por AdminOnly.
//
// VERIFICAR:
// Debe responder 403 Forbidden.
// =======================================
func TestAdminOnly_RolCliente(t *testing.T) {

	// Handler protegido para ADMIN
	handler := AdminOnly(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Request de prueba
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Simular rol CLIENTE en el contexto
	ctx := context.WithValue(req.Context(), RolKey, "CLIENTE")
	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	// Ejecutar
	handler.ServeHTTP(rec, req)

	// Verificar
	if rec.Code != http.StatusForbidden {
		t.Errorf(
			"se esperaba %d pero llegó %d",
			http.StatusForbidden,
			rec.Code,
		)
	}
}