package models

import (
	"encoding/hex"
	"crypto/md5"
)

type User struct{
	id int
	firstName string
	lastName string
	email string
	password string
}

type Login []User

func NewLogin()(Login){
	return make(Login, 0)
}

func (l *Login) Add(firstName string, lastName string, email string, password string) User{
	user:=User{
		id: IncrementUserId(*l),
		firstName: firstName,
		lastName: lastName,
		email: email,
		password: GetMD5Hash(password),
	}

	*l = append(*l, user)
	return user
	
}

func (l *Login) GetUserByEmail(email string) string{
	return "abc"
}

func IncrementUserId(l Login) int{
	return len(l) + 1
}

func GetMD5Hash(str string)string{
	hasher := md5.New()
    hasher.Write([]byte(str))
    return hex.EncodeToString(hasher.Sum(nil))
}



