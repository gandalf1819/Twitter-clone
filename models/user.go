package models

import (
	"crypto/md5"
	"encoding/hex"
	//"log"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Follows   []int
}

//userType can be either Follower or Unfollower
type UserList struct {
	Id        int
	FirstName string
	LastName  string
	UserType  string
}

type Login []User

func NewLogin() Login {
	return make(Login, 0)
}

func (l *Login) Add(firstName string, lastName string, email string, password string) User {
	user := User{
		Id:        IncrementUserId(*l),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  GetMD5Hash(password),
		Follows:   make([]int, 0),
	}

	*l = append(*l, user)
	return user

}

func (l *Login) GetUserByEmailPassword(email string, password string) User {
	var userObj User
	for id, value := range *l {
		if value.Email == email && value.Password == GetMD5Hash(password) {
			userObj = (*l)[id]
		}
	}
	return userObj
}

func (l *Login) FollowUser(userId int, followerId int) {

	for id, value := range *l {
		if value.Id == userId {
			(*l)[id].Follows = append((*l)[id].Follows, followerId)
			break
		}
	}

}

func (l *Login) UnfollowUser(userId int, followerId int) {
	for id, value := range *l {
		if value.Id == userId {
			length := len((*l)[id].Follows) - 1
			for index, currentValue := range (*l)[id].Follows {
				if currentValue == followerId {
					(*l)[id].Follows[index], (*l)[id].Follows[length] = (*l)[id].Follows[length], (*l)[id].Follows[id]
					break
				}
			}

			(*l)[id].Follows = (*l)[id].Follows[:length]
			return
		}
	}
}

func (l *Login) GetFollowerSuggestions(userId int) []UserList {

	var userObj User
	userList := make([]UserList, 0)

	for _, user := range *l {
		if user.Id == userId {
			userObj = user
		} else {
			currentUser := UserList{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				UserType:  "Unfollower",
			}
			userList = append(userList, currentUser)
		}
	}

	for _, userId := range userObj.Follows {
		for id, userListObj := range userList {
			if userId == userListObj.Id {
				userList[id].UserType = "Follower"
				break
			}
		}
	}

	return userList

}

func IncrementUserId(l Login) int {
	return len(l) + 1
}

func GetMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
