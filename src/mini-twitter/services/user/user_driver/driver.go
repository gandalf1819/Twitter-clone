package user_driver

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"log"
	"mini-twitter/services/user/userpb"
	"mini-twitter/util"
)

type Server struct{}

var lo userpb.Login

func Init() {
	_, err := util.InteractWithRaftStorage("PUT", "userDB", lo)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}

	log.Println("DB User Initialized =", lo.Users)
}

func GetUserDB(value interface{}) (userpb.Login, error) {
	var db userpb.Login
	data, err := util.InteractWithRaftStorage("GET", "userDB", db)
	if err != nil {
		log.Println("Error occured while getting user data from Raft =", err)
		panic(err)
	}
	var userDB userpb.Login
	userDB, err = DecodeRaftUserStorage(data)
	if err != nil {
		log.Println("Error occured while decoding user data from Raft storage =", err)
		return userDB, err
	}
	log.Println("userDB after decode =", userDB)
	return userDB, nil
}

func DecodeRaftUserStorage(db string) (userpb.Login, error) {
	log.Println("Decode User Storage called")
	dec := gob.NewDecoder(bytes.NewBufferString(db))
	if err := dec.Decode(&lo); err != nil {
		log.Fatalf("raftexample: could not decode message (%v)", err)
		return lo, err
	}
	//log.Println("userDB in DecodeRaftTokenStorage =", lo)

	return lo, nil
}

func (*Server) Add(ctx context.Context, userParams *userpb.AddUserParameters) (*userpb.User, error) {
	var lo userpb.Login
	user := &userpb.User{
		Id:        IncrementUserId(),
		FirstName: userParams.FirstName,
		LastName:  userParams.LastName,
		Email:     userParams.Email,
		Password:  GetMD5Hash(userParams.Password),
		Follows:   make([]int32, 0),
	}

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}

	for _, value := range userDB.Users {
		if value.Email == user.Email {
			return nil, errors.New("User already registered!!")
		}
	}

	userDB.Users = append(userDB.Users, user)

	_, err = util.InteractWithRaftStorage("PUT", "userDB", userDB)
	if err != nil {
		log.Println("Error occured while storing user data in Raft =", err)
		panic(err)
	}
	log.Println("User added =", user)
	log.Println("User DB =", userDB.Users)
	return user, nil
}

func (*Server) GetUserByEmailPassword(ctx context.Context, loginParams *userpb.LoginDetails) (*userpb.User, error) {
	var lo userpb.Login

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}

	var userObj userpb.User
	for id, value := range userDB.Users {
		if value.Email == loginParams.Email && value.Password == GetMD5Hash(loginParams.Password) {
			userObj = *userDB.Users[id]
		}
	}
	return &userObj, nil
}

func (*Server) FollowUser(ctx context.Context, fp *userpb.FollowerParameters) (*userpb.Status, error) {
	var lo userpb.Login

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}

	for id, value := range userDB.Users {
		if value.Id == fp.UserId {
			userDB.Users[id].Follows = append(userDB.Users[id].Follows, fp.FollowerId)
			break
		}
	}
	_, err = util.InteractWithRaftStorage("PUT", "userDB", userDB)
	if err != nil {
		log.Println("Error occured while storing user data in Raft =", err)
		panic(err)
	}
	log.Println("User ", fp.UserId, " follows ", fp.FollowerId)
	log.Println("User DB =", userDB.Users)

	return &userpb.Status{ResponseStatus: true}, nil
}

func (*Server) UnfollowUser(ctx context.Context, fp *userpb.FollowerParameters) (*userpb.Status, error) {
	log.Println("UnfollowUser called =")
	var lo userpb.Login

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}
	for id, value := range userDB.Users {
		if value.Id == fp.UserId {
			log.Println("value =", value)
			length := len(userDB.Users[id].Follows) - 1
			log.Println("length =", length)
			for index, currentValue := range userDB.Users[id].Follows {
				if currentValue == fp.FollowerId {
					userDB.Users[id].Follows[index], userDB.Users[id].Follows[length] = userDB.Users[id].Follows[length], userDB.Users[id].Follows[index]
					break
				}
			}

			userDB.Users[id].Follows = userDB.Users[id].Follows[:length]

			_, err = util.InteractWithRaftStorage("PUT", "userDB", userDB)
			if err != nil {
				log.Println("Error occured while storing user data in Raft =", err)
				panic(err)
			}
			log.Println("UserId ", fp.UserId, " unfollows ", fp.FollowerId)
			log.Println("User DB =", userDB.Users)

			return &userpb.Status{ResponseStatus: true}, nil
		}
	}
	log.Println("UserId ", fp.UserId, " not found to unfollow ", fp.FollowerId, " user")
	return &userpb.Status{ResponseStatus: false}, nil
}

func (*Server) GetFollowerSuggestions(ctx context.Context, userId *userpb.UserId) (*userpb.UserList, error) {
	log.Println("Get Follower Suggestions called")

	var lo userpb.Login

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}

	log.Println("userDB after calling Get Follower Suggestions =", userDB)

	var userObj *userpb.User
	userList := userpb.UserList{
		List: make([]*userpb.UserListFields, 0),
	}

	for _, user := range userDB.Users {
		if user.Id == userId.Id {
			userObj = user
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
	var lo userpb.Login

	userDB, err := GetUserDB(lo)
	if err != nil {
		return nil, err
	}
	var userObj *userpb.User
	userListObj := &userpb.Login{
		Users: make([]*userpb.User, 0),
	}

	for id, value := range userDB.Users {
		if value.Id == user.Id {
			userObj = userDB.Users[id]
			break
		}
	}

	userListObj.Users = append(userListObj.Users, userObj)

	for _, value := range userObj.Follows {
		for _, user := range userDB.Users {
			if value == user.Id {
				userListObj.Users = append(userListObj.Users, user)
				break
			}
		}
	}

	return userListObj, nil
}

func (*Server) GetAllUsers(ctx context.Context, in *userpb.NoArgs) (*userpb.Login, error) {
	var err error
	log.Println("GetAllUsers called")
	lo, err = GetUserDB(lo)
	if err != nil {
		return nil, err
	}
	return &lo, nil
}

func GetMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IncrementUserId() int32 {
	var lo userpb.Login

	userDB, _ := GetUserDB(lo)

	return int32(len(userDB.Users) + 1)
}
