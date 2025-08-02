package service

import (
	"errors"
	"fmt"

	"gabyconta/usuarios/model"
	"gorm.io/gorm"
)

type UsuarioService struct {
	db *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

func (s *UsuarioService) CrearUsuario(usuario *model.Usuario) error {
	// Aquí deberías hashear la contraseña antes de guardar (bcrypt)
	return s.db.Create(usuario).Error
}

func (s *UsuarioService) Autenticar(email, password string) (*model.Usuario, error) {
	var usuario model.Usuario
	if err := s.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	// Aquí valida el password con bcrypt.CompareHashAndPassword (ejemplo simple)
	if usuario.Password != password {
		return nil, errors.New("contraseña incorrecta")
	}

	return &usuario, nil
}

func (s *UsuarioService) CargarEmpresas(usuario *model.Usuario) error {
	return s.db.Preload("Empresas").First(usuario, usuario.ID).Error
}

func (s *UsuarioService) GenerarToken(userID, empresaID uint) (string, error) {
	// Aquí deberías implementar JWT real con firma, expiración, etc.
	token := fmt.Sprintf("token-usuario-%d-empresa-%d", userID, empresaID)
	return token, nil
}
