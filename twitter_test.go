package handler

import (
	"log"
	"testing"

	"./handler"
	"./models"
)

type database struct {
	l  models.Login
	t  models.Token
	up models.UserPosts
}

var db database

func createTestData() {
	handler.Init()
}

func TestRegister(t *testing.T) {
	log.Println("EXECUTING TestRegister TEST CASE")
	createTestData()
	var registerData []models.User
	registerData = []models.User{
		models.User{
			Id:        0,
			FirstName: "Nikhil",
			LastName:  "Nar",
			Email:     "ncn251@nyu.edu",
			Password:  "test",
			Follows:   make([]int, 0),
		},
		models.User{
			Id:        0,
			FirstName: "Chinmay",
			LastName:  "Wyawahare",
			Email:     "cnw282@nyu.edu",
			Password:  "test",
			Follows:   make([]int, 0),
		},
		models.User{
			Id:        0,
			FirstName: "Yuvraj",
			LastName:  "Singh",
			Email:     "ys102@nyu.edu",
			Password:  "test",
			Follows:   make([]int, 0),
		},
		models.User{
			Id:        0,
			FirstName: "Gahan",
			LastName:  "Jagadeesh",
			Email:     "gj100@nyu.edu",
			Password:  "test",
			Follows:   make([]int, 0),
		},
	}

	for id, value := range registerData {
		db.l.Add(value.FirstName, value.LastName, value.Email, value.Password)
		if len(db.l) != (id + 1) {
			t.Errorf("Data with first name %v and last name %v not inserted", value.FirstName, value.LastName)
		}
	}
	log.Println("After test case 1, Login Object =", db.l)
}

func TestFollowUser(t *testing.T) {
	log.Println("EXECUTING TestFollowUser TEST CASE")
	followerMap := make(map[string]int)
	//Nikhil Follows Yuvraj
	db.l.FollowUser(1, 3)
	followerMap["Yuvraj"] = 0
	//Nikhil Follows Chinmay
	db.l.FollowUser(1, 2)
	followerMap["Chinmay"] = 0
	users := db.l.GetUserFollowersById(1)

	for _, value := range users {
		followerMap[value.FirstName] = 1
	}

	for key, value := range followerMap {
		if value != 1 {
			t.Errorf("Nikhil failed to follow %v", key)
		}

	}
	log.Println("After test case 2, Login Object =", db.l)
}

func TestUnFollowUser(t *testing.T) {
	log.Println("EXECUTING TestUnFollowUser TEST CASE")
	unFollowerMap := make(map[string]int)
	//Nikhil UnFollows Yuvraj
	db.l.UnfollowUser(1, 3)
	unFollowerMap["Yuvraj"] = 0
	users := db.l.GetUserFollowersById(1)

	for _, value := range users {
		if value.Id == 3 {
			unFollowerMap[value.FirstName] = 1
		}
	}

	for key, value := range unFollowerMap {
		if value != 0 {
			t.Errorf("Nikhil failed to unfollow %v", key)
		}
	}
	log.Println("After test case 3, Login Object =", db.l)

}

func TestPosts(t *testing.T) {
	log.Println("EXECUTING TestPosts TEST CASE")
	postsMap := make(map[string]int)
	//Yuvraj adds status
	db.up.AddPost(3, "This is a Yuvraj's status")
	postsMap["Yuvraj"] = 0
	//Nikhil adds status
	db.up.AddPost(1, "This is a Nikhil's status")
	postsMap["Nikhil"] = 0
	//Chinmay adds status
	db.up.AddPost(2, "This is a Chinmay's status")
	postsMap["Chinmay"] = 0

	for _, value := range db.up {
		if value.UserId == 1 {
			postsMap["Nikhil"] = 1
		} else if value.UserId == 2 {
			postsMap["Chinmay"] = 1
		} else if value.UserId == 3 {
			postsMap["Yuvraj"] = 1
		}
	}

	for key, value := range postsMap {
		if value != 1 {
			t.Errorf("%v failed to add post", key)
		}

	}
	log.Println("After test case 4, Login Object =", db.l)
	log.Println("After test case 4, Posts Object =", db.up)

}

func TestGetFollowerPosts(t *testing.T) {
	log.Println("EXECUTING TestGetFollowerPosts TEST CASE")
	postsMap := make(map[string]int)

	// Get Nikhil's newsfeed
	nikPosts := db.l.GetFollowerPosts(1, &db.up)
	postsMap["Nikhil"] = 0
	postsMap["Chinmay"] = 0

	for _, value := range nikPosts {
		if value.Id == 1 {
			postsMap["Nikhil"] = 1
		} else if value.Id == 2 {
			postsMap["Chinmay"] = 1
		}
	}

	for key, value := range postsMap {
		if value != 1 {
			t.Errorf("%v post was not found in Nikhil's newsfeed", key)
		}

	}
	log.Println("After test case 5, Login Object =", db.l)
	log.Println("After test case 5, Posts Object =", db.up)

}
