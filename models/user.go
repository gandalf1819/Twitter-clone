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
	follows []int
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
		follows:make([]int,0),
	}

	*l = append(*l, user)
	return user
	
}

func (l *Login) GetUserByEmail(email string) string{
	return "abc"
}

func (l *Login) FollowUser(userId int, followerId int){
	var user User
	for _,value := range *l{
		if value.id == userId{
			user = value
			break 
		}
	}

	user.follows = append(user.follows, followerId)
}

func (l *Login) UnfollowUser(userId int, followerId int){
	var user User
	for _,value := range *l{
		if value.id == userId{
			user = value
			break 
		}
	}
	length:= len(user.follows)-1
	for id, value := range user.follows{
		if value == followerId{
			user.follows[id], user.follows[length] = user.follows[length], user.follows[id]
			break
		}
	} 
		
	length= length -1	
	user.follows = user.follows[:length]
}

func IncrementUserId(l Login) int{
	return len(l) + 1
}

func GetMD5Hash(str string)string{
	hasher := md5.New()
    hasher.Write([]byte(str))
    return hex.EncodeToString(hasher.Sum(nil))
}



