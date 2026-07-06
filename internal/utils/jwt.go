package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("clave-secreta-barberia")

func GenerarToken(usuarioID int, correo string, rol string) (string, error) {
	claims := jwt.MapClaims{
		"usuario_id": usuarioID,
		"correo":     correo,
		"rol":        rol,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}