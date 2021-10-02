package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID int, roleID int, nama string, email string, gambarUser string, telp string, alamat string) string
	ValidateToken(token string) (*jwt.Token, error)
}

// type jwtCustomClaim struct {
// 	UserID string `json:"user_id"`
// 	jwt.StandardClaims
// }

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "adhiardiansyah",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "adhiardiansyah"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userID int, roleID int, nama string, email string, gambarUser string, telp string, alamat string) string {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["role_id"] = roleID
	claims["nama"] = nama
	claims["email"] = email
	claims["gambar_user"] = gambarUser
	claims["telp"] = telp
	claims["alamat"] = alamat
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode masuk yang tidak terduga %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
