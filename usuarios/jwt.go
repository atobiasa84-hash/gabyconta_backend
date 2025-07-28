package usuarios

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gabyconta/usuarios/model"
)

// jwtKey es la clave secreta usada para firmar los tokens JWT.
// En producción, usa una variable de entorno segura.
var jwtKey = []byte(getJWTSecret())

// getJWTSecret obtiene la clave secreta desde variable de entorno o valor por defecto
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "superuser" // Valor por defecto si no se configura la variable
	}
	return secret
}

// Claims define los datos que incluirá el token JWT
type Claims struct {
	UsuarioID uint   `json:"usuario_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

// GenerarToken crea un token JWT firmado válido por 24 horas
func GenerarToken(usuario *model.Usuario) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UsuarioID: usuario.ID,
		Email:     usuario.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gabyconta",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


