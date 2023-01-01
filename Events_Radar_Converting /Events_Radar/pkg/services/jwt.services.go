package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"radar/pkg/model"
	service "radar/pkg/services/interface"
)




type jwtService struct {
	SecretKey string
}

func NewJWTUserService() service.JWTService {
	return &jwtService{
		SecretKey: os.Getenv("USER_KEY"),
	}
}



func NewJWTAdminService() service.JWTService {
	return &jwtService{
		SecretKey: os.Getenv("ADMIN_KEY"),
	}
}

func (j *jwtService) GenerateToken(userId int, email, role string) string {
	claims := &model.SignedDetails{
		User_Id :userId,
		Username :email,
		Role:role,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(2)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken
}
func (j *jwtService) GenerateRefreshToken(accessToken string) (string, error) {
	
	claims := &model.SignedDetails{}
	j.GetTokenFromString(accessToken, claims)

	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		return "", errors.New("too early to generate refresh token")
	}

	claims.ExpiresAt = time.Now().Local().Add(time.Minute * time.Duration(5)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}
	return refreshToken, err

}

func (j *jwtService) VerifyToken(signedToken string) (bool, *model.SignedDetails) {

	claims := &model.SignedDetails{}
	token, _ := j.GetTokenFromString(signedToken, claims)

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func (j *jwtService) GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error) {

	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}
