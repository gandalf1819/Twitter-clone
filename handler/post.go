package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../models"
)

type PostsPageData struct {
	Friends []models.UserList
}

type PostData struct {
	UserId int
	Text   string
}

type Follow struct {
	UserId     int
	FollowerId int
}

func Posts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		cookieToken, err := r.Cookie("token")
		if err != nil || cookieToken.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		userId, err := db.t.GetUserIdFromToken(cookieToken.Value)
		if err != nil {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		postsData := PostsPageData{
			Friends: db.l.GetFollowerSuggestions(userId),
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
		ReturnAPIResponse(w, r, 200, "Status shared successfully!!", make(map[string]string))
	}
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}
		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		follower := followsData.FollowerId
		db.l.FollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User Followed successfully!!", make(map[string]string))
	}
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}
		var followsData Follow
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &followsData)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		follower := followsData.FollowerId
		db.l.UnfollowUser(user, follower)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User UnFollowed successfully!!", make(map[string]string))
	}
}
