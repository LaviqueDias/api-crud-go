package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type User struct {
	ID        string
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
