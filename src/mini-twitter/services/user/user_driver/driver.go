package user_driver

import (
	"../userpb"
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
)

type Server struct{}

var lo userpb.Login

func Init() {
	lo := NewLogin()

	log.Println("DB User Initialized =", lo.Users)
}

func NewLogin() *userpb.Login {
	userPosts := &userpb.Login{
		Users: make([]*userpb.User, 0),
	}
	return userPosts
}

func (*Server) Add(ctx context.Context, userParams *userpb.AddUserParameters) (*userpb.User, error) {
	user := &userpb.User{
		Id:        IncrementUserId(),
		FirstName: userParams.FirstName,
		LastName:  userParams.LastName,
		Email:     userParams.Email,
		Password:  GetMD5Hash(userParams.Password),
		Follows:   make([]int32, 0),
	}

	lo.Users = append(lo.Users, user)
	log.Println("User added =", user)
	log.Println("User DB =", lo.Users)
	return user, nil
}

func (*Server) GetUserByEmailPassword(ctx context.Context, loginParams *userpb.LoginDetails) (*userpb.User, error) {
	var userObj userpb.User
	for id, value := range lo.Users {
		if value.Email == loginParams.Email && value.Password == GetMD5Hash(loginParams.Password) {
			userObj = *lo.Users[id]
		}
	}
	return &userObj, nil
}

func (*Server) FollowUser(ctx context.Context, fp *userpb.FollowerParameters) (*userpb.Status, error) {
	for id, value := range lo.Users {
		if value.Id == fp.UserId {
			lo.Users[id].Follows = append(lo.Users[id].Follows, fp.FollowerId)
			break
		}
	}
	log.Println("User ", fp.UserId, " follows ", fp.FollowerId)
	log.Println("User DB =", lo.Users)

	return &userpb.Status{ResponseStatus: true}, nil
}

func (*Server) UnfollowUser(ctx context.Context, fp *userpb.FollowerParameters) (*userpb.Status, error) {
	log.Println("UnfollowUser called =")
	for id, value := range lo.Users {
		if value.Id == fp.UserId {
			log.Println("value =", value)
			length := len(lo.Users[id].Follows) - 1
			log.Println("length =", length)
			for index, currentValue := range lo.Users[id].Follows {
				if currentValue == fp.FollowerId {
					lo.Users[id].Follows[index], lo.Users[id].Follows[length] = lo.Users[id].Follows[length], lo.Users[id].Follows[index]
					break
				}
			}

			lo.Users[id].Follows = lo.Users[id].Follows[:length]
			log.Println("UserId ", fp.UserId, " unfollows ", fp.FollowerId)
			log.Println("User DB =", lo.Users)

			return &userpb.Status{ResponseStatus: true}, nil
		}
	}
	log.Println("UserId ", fp.UserId, " not found to unfollow ", fp.FollowerId, " user")
	return &userpb.Status{ResponseStatus: false}, nil
}

func (*Server) GetFollowerSuggestions(ctx context.Context, userId *userpb.UserId) (*userpb.UserList, error) {
	var userObj userpb.User
	userList := userpb.UserList{
		List: make([]*userpb.UserListFields, 0),
	}

	for _, user := range lo.Users {
		if user.Id == userId.Id {
			userObj = *user
		} else {
			currentUser := &userpb.UserListFields{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				UserType:  "Unfollower",
			}
			userList.List = append(userList.List, currentUser)
		}
	}

	for _, userId := range userObj.Follows {
		for id, userListObj := range userList.List {
			if userId == userListObj.Id {
				userList.List[id].UserType = "Follower"
				break
			}
		}
	}

	log.Println("Follower Suggestions for userId ", userId.Id, " =", userList.List)

	return &userList, nil
}

func (*Server) GetUserFollowersById(ctx context.Context, user *userpb.UserId) (*userpb.Login, error) {
	var userObj *userpb.User
	userListObj := &userpb.Login{
		Users: make([]*userpb.User, 0),
	}

	for id, value := range lo.Users {
		if value.Id == user.Id {
			userObj = lo.Users[id]
			break
		}
	}

	userListObj.Users = append(userListObj.Users, userObj)

	for _, value := range userObj.Follows {
		for _, user := range lo.Users {
			if value == user.Id {
				userListObj.Users = append(userListObj.Users, user)
				break
			}
		}
	}

	return userListObj, nil
}

func (*Server) GetAllUsers(ctx context.Context, in *userpb.NoArgs) (*userpb.Login, error) {
	return &lo, nil
}

func GetMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IncrementUserId() int32 {
	return int32(len(lo.Users) + 1)
}
