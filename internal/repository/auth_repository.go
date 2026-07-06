package repository

import (
	"barberia-cola-virtual/internal/models"
	"barberia-cola-virtual/internal/storage"
)

func CreateUsuario(usuario models.Usuario) models.Usuario {
	usuario.ID = len(storage.Usuarios) + 1
	storage.Usuarios = append(storage.Usuarios, usuario)
	return usuario
}

func GetUsuarioByCorreo(correo string) (models.Usuario, bool) {
	for _, usuario := range storage.Usuarios {
		if usuario.Correo == correo {
			return usuario, true
		}
	}

	return models.Usuario{}, false
}