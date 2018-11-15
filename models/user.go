package models

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	password  string
	follows   []int
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
		id:        IncrementUserId(*l),
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  GetMD5Hash(password),
		follows:   make([]int, 0),
	}

	*l = append(*l, user)
	return user

}

func (l *Login) GetUserByEmail(email string) string {
	return "abc"
}

func (l *Login) FollowUser(userId int, followerId int) {

	for id, value := range *l {
		if value.id == userId {
			(*l)[id].follows = append((*l)[id].follows, followerId)
			break
		}
	}

}

func (l *Login) UnfollowUser(userId int, followerId int) {
	for id, value := range *l {
		if value.id == userId {
			length := len((*l)[id].follows) - 1
			for index, currentValue := range (*l)[id].follows {
				if currentValue == followerId {
					(*l)[id].follows[index], (*l)[id].follows[length] = (*l)[id].follows[length], (*l)[id].follows[id]
					break
				}
			}

			(*l)[id].follows = (*l)[id].follows[:length]
			return
		}
	}
}

func (l *Login) GetFollowerSuggestions(userId int) []UserList {

	var userObj User
	userList := make([]UserList, 0)

	for _, user := range *l {
		if user.id == userId {
			userObj = user
		} else {
			currentUser := UserList{
				Id:        user.id,
				FirstName: user.firstName,
				LastName:  user.lastName,
				UserType:  "Unfollower",
			}
			userList = append(userList, currentUser)
		}
	}

	for _, userId := range userObj.follows {
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
