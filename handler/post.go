package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"../models"
)

type PostData struct {
	UserId int
	Text   string
}

type PostsPageData struct {
	Friends []models.UserList
}

type Follow struct {
	UserId     int
	FollowerId int
}

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		postsData := PostsPageData{

			Friends: db.l.GetFollowerSuggestions(1),
		}

		t, _ := template.ParseFiles("./views/html/posts.html")
		t.Execute(w, postsData)

	} else {
		r.ParseForm()

		var post PostData
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &post)
		user := post.UserId
		status := post.Text
		//used to post data

		db.up.AddPost(user, status)
		log.Println("db.up===", db.up)
		ReturnAPIResponse(w, r, 200, "Status shared successfully!!")
	}
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user := followsData.UserId
		follower := followsData.FollowerId
		db.l.FollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User Followed successfully!!")
	}
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user := followsData.UserId
		follower := followsData.FollowerId
		db.l.UnfollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User UnFollowed successfully!!")
	}
}
