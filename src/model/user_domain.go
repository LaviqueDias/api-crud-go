package model

import (
	"crypto/md5"
	"encoding/hex"

)

type UserDomainInterface interface {
    GetEmail() string
    GetPassword() string
    GetName() string
    GetRole() string
    IsActive() bool

    EncryptPassword()
}
type userDomain struct {
    name      string    
    email     string    
    password  string    
    role      string    
    active    bool      
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetRole() string {
	return ud.role
}

func (ud *userDomain) IsActive() bool {
	return ud.active
}

func NewUserDomain(name, email, password, role string, active bool) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		role:     role,
		active:   active,
	}
}

func(ud *userDomain) EncryptPassword(){
    hash := md5.New()
    defer hash.Reset()
    hash.Write([]byte(ud.password))
    ud.password = hex.EncodeToString(hash.Sum(nil))
}

