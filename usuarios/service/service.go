package service

import (
	"errors"

	"gabyconta/usuarios/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsuarioService struct {
	DB *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

func (s *UsuarioService) CrearUsuario(usuario *model.Usuario) error {
	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuario.Password = string(hashedPassword)

	return s.DB.Create(usuario).Error
}

func (s *UsuarioService) Autenticar(email, password string) (*model.Usuario, error) {
	var usuario model.Usuario
	if err := s.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password)); err != nil {
		return nil, errors.New("contraseña incorrecta")
	}

	return &usuario, nil
}