package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
)

type UserDomain struct {
    Name      string    
    Email     string    
    Password  string    
    Role      string    
    Active    bool      
}

func(ud *UserDomain) EncryptPassword(){
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
    CreateUser() *rest_err.RestErr
    UpdateUser(string) *rest_err.RestErr
    FindUser(string) (UserDomain, *rest_err.RestErr)
    DeleteUser(string) *rest_err.RestErr
}