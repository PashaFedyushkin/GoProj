package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	project "github.com/PashaFedyushkin/GoProj"
	"github.com/PashaFedyushkin/GoProj/pkg/repos"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt    = "qwerty"
	signKey = "qazwsxedc"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	rep repos.Authorizaton
}

func NewAuthService(rep repos.Authorizaton) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) CreateUser(user project.User) (int, error) {
	user.Password = genPassHash(user.Password)
	return s.rep.CreateUser(user)
}

func genPassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.rep.GetUser(username, genPassHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signKey))
}
