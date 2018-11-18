package handler

import (
	"log"
	"testing"

	"../mini-twitter/models"
)

// database initialization
type database struct {
	l  models.Login
	t  models.Token
	up models.UserPosts
}

var db database

func Init() {
	l := models.NewLogin()
	t := models.NewToken()
	up := models.NewUserPosts()
	db = database{l, t, up}
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//userType can be either Follower or Unfollower
type UserList struct {
	Id        int
	FirstName string
	LastName  string
	UserType  string
}

type Login []User

func addTestUserData(FirstName string, LastName string, Email string, Password string) {

	db.l.Add(FirstName, LastName, Email, Password)
	log.Println("db.l===", db.l)
}

func TestLogin(t *testing.T) {

	// add two test entries
	addTestUserData("Chinmay", "Wyawahare", "xyz@gmail.com", "qwerty")
	addTestUserData("Nikhil", "Nar", "abc@gmail.com", "zxcvb")

}

func TestPosts(t *testing.T) {

	db.up.AddPost(999, "This is a test status")
	log.Println("db.up===", db.up)
}

func TestFollowUser(t *testing.T) {

	db.l.FollowUser(9, 99)
	log.Println("db.l===", db.l)
}

func TestUnfollowUser(t *testing.T) {

	db.l.UnfollowUser(99, 9)
	log.Println("db.l===", db.l)
}
