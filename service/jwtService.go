package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GeneratedToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIdByToken(tokenStr string) string
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "chat-system",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "chat-system"
	}
	return secretKey
}

// GeneratedToken 生成Token
func (j *jwtService) GeneratedToken(userId string) string {
	claims := jwtCustomClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
			Issuer:    j.secretKey,
			IssuedAt:  time.Now().Unix(),
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := withClaims.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return token
}

// ValidateToken 验证Token
func (j *jwtService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserIdByToken(tokenStr string) string {
	token, err := j.ValidateToken(tokenStr)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
