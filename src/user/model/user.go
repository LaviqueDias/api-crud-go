package model

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"

	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Role      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func (u *User) ComparePassword(password string) bool {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)) == u.Password
}

func (u *User) GenerateToken() (string, *rest_err.RestErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		// logger.Error("Error signing JWT token", err,
		// 	zap.String("journey", "loginUser"),
		// )
		return "", rest_err.NewInternalServerError("error generating token")
	}

	return tokenString, nil
}
