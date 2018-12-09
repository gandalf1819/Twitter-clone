package test

import (
	"../userpb"
	"context"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	log.Println("EXECUTING TestRegister TEST CASE")
	InitializeUserClient()
	var registerData []userpb.AddUserParameters
	registerData = []userpb.AddUserParameters{
		userpb.AddUserParameters{
			FirstName: "Nikhil",
			LastName:  "Nar",
			Email:     "ncn251@nyu.edu",
			Password:  "test",
		},
		userpb.AddUserParameters{
			FirstName: "Chinmay",
			LastName:  "Wyawahare",
			Email:     "cnw282@nyu.edu",
			Password:  "test",
		},
		userpb.AddUserParameters{
			FirstName: "Yuvraj",
			LastName:  "Singh",
			Email:     "ys102@nyu.edu",
			Password:  "test",
		},
		userpb.AddUserParameters{
			FirstName: "Gahan",
			LastName:  "Jagadeesh",
			Email:     "gj100@nyu.edu",
			Password:  "test",
		},
	}

	for id, value := range registerData {

		_, err := uc.UserDB.Add(context.Background(), &value)
		if err != nil {
			t.Errorf("Data with first name %v and last name %v not inserted", value.FirstName, value.LastName)
			log.Println("Error received from User Service =", err)
			return
		}

		var allUsers *userpb.Login
		allUsers, err = uc.UserDB.GetAllUsers(context.Background(), &userpb.NoArgs{})
		if err != nil {
			t.Errorf("Data with first name %v and last name %v not inserted", value.FirstName, value.LastName)
			log.Println("Error received from User Service =", err)
			return
		}

		if len(allUsers.Users) != (id + 1) {
			t.Errorf("Data with first name %v and last name %v not inserted", value.FirstName, value.LastName)
		}
	}
	log.Println("TestRegister Test Passed")
}

func TestFollowUser(t *testing.T) {
	log.Println("EXECUTING TestFollowUser TEST CASE")
	followerMap := make(map[string]int)
	//Nikhil Follows Yuvraj

	fp := &userpb.FollowerParameters{
		UserId:     int32(1),
		FollowerId: int32(3),
	}
	_, err := uc.UserDB.FollowUser(context.Background(), fp)
	if err != nil {
		log.Println("Error received from User Service =", err)
		return
	}

	followerMap["Yuvraj"] = 0
	//Nikhil Follows Chinmay
	fp = &userpb.FollowerParameters{
		UserId:     int32(1),
		FollowerId: int32(2),
	}
	_, err = uc.UserDB.FollowUser(context.Background(), fp)
	if err != nil {
		log.Println("Error received from User Service =", err)
		return
	}
	followerMap["Chinmay"] = 0

	user := &userpb.UserId{
		Id: int32(1),
	}

	var users *userpb.Login
	users, err = uc.UserDB.GetUserFollowersById(context.Background(), user)

	for _, value := range users.Users {
		followerMap[value.FirstName] = 1
	}

	for key, value := range followerMap {
		if value != 1 {
			t.Errorf("Nikhil failed to follow %v", key)
		}

	}
	log.Println("TestFollowUser Test Passed")
}

func TestUnFollowUser(t *testing.T) {
	log.Println("EXECUTING TestUnFollowUser TEST CASE")
	unFollowerMap := make(map[string]int)
	//Nikhil UnFollows Yuvraj
	fp := &userpb.FollowerParameters{
		UserId:     int32(1),
		FollowerId: int32(3),
	}
	_, err := uc.UserDB.UnfollowUser(context.Background(), fp)
	if err != nil {
		log.Println("Error received from User Service =", err)
		return
	}

	unFollowerMap["Yuvraj"] = 0

	user := &userpb.UserId{
		Id: int32(1),
	}
	var users *userpb.Login
	users, err = uc.UserDB.GetUserFollowersById(context.Background(), user)

	for _, value := range users.Users {
		if value.Id == 3 {
			unFollowerMap[value.FirstName] = 1
		}
	}

	for key, value := range unFollowerMap {
		if value != 0 {
			t.Errorf("Nikhil failed to unfollow %v", key)
		}
	}
	log.Println("TestUnFollowUser Test Passed")

}
