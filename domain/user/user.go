package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/dannegm/spendzer-service/config"
)

var (
	ErrInvalidUser = errors.New("Is not a valid User")
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	JWT      string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Security.SALT)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func New(username string, password string) (User, error) {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return User{}, err
	}

	return User{
		ID:       uuid.New(),
		Username: username,
		Password: hashedPassword,
	}, nil
}

func (u *User) setJWT(jwt string) error {
	if u == nil {
		return errors.New("Provide a valid user")
	}

	u.JWT = jwt

	return nil
}

func (u *User) getJWT() string {
	return u.JWT
}

func (u *User) generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token.SignedString(config.App.Secret)
}
