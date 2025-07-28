package model

type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

// LoginInput define los campos esperados para el login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
