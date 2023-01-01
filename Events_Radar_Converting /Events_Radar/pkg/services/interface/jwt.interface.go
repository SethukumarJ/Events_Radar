package interfaces

import (
	"radar/pkg/model"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(user_id int, email string, role string) string
	VerifyToken(token string) (bool, *model.SignedDetails)
	GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error)
	GenerateRefreshToken(token string) (string, error)
}
