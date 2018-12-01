package handler

import (
	"../models"
	"../services/auth/authpb"
	"../services/post/postpb"
	"../services/user/userpb"
	"context"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PostsPageData struct {
	Friends []*userpb.UserListFields
	Posts   []models.PostsList
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

		tokName := &authpb.AuthTokenName{
			TokenName: cookieToken.Value,
		}

		userId, err := con.GetAuthTokenClient().GetUserIdFromToken(context.Background(), tokName)
		if err != nil {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			log.Println("Error received from Auth Service =", err)
			return
		}

		user := &userpb.UserId{
			Id: int32(userId.TokenValue),
		}

		friendLists, err := con.GetUserClient().GetFollowerSuggestions(context.Background(), user)
		if err != nil {
			log.Println("Error received from User Service =", err)
			friendLists = &userpb.UserList{
				List: make([]*userpb.UserListFields, 0),
			}
		}

		postsData := PostsPageData{
			Friends: friendLists.List,
			Posts:   db.l.GetFollowerPosts(int(userId.TokenValue), &db.up),
		}
		log.Println("Posts=======", postsData.Posts)

		t, _ := template.ParseFiles("./views/html/posts.html")
		t.Execute(w, postsData)

	} else if r.Method == "POST" {
		type Status struct {
			Status string
		}
		var statusMessage Status

		userCookie, err := r.Cookie("user_id")
		if err != nil || userCookie.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &statusMessage)
		user, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			panic(err)
		}
		text := statusMessage.Status

		db.up.AddPost(user, text)
		userPost := &postpb.PostText{
			UserId: int32(user),
			Text:   text,
		}
		_, err = con.GetUserPostClient().AddPost(context.Background(), userPost)
		if err != nil {
			ReturnAPIResponse(w, r, 422, "Error occured while adding post. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from UserPost Service =", err)
			return
		}
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
		//db.l.FollowUser(user, follower)
		fp := &userpb.FollowerParameters{
			UserId:     int32(user),
			FollowerId: int32(follower),
		}
		_, err = con.GetUserClient().FollowUser(context.Background(), fp)
		if err != nil {
			ReturnAPIResponse(w, r, 422, "Error occured while following a friend. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from User Service =", err)
			return
		}
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
		//db.l.UnfollowUser(user, follower)
		fp := &userpb.FollowerParameters{
			UserId:     int32(user),
			FollowerId: int32(follower),
		}
		status, err := con.GetUserClient().UnfollowUser(context.Background(), fp)
		if err != nil || !status.ResponseStatus {
			ReturnAPIResponse(w, r, 422, "Error occured while unfollowing a friend. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from User Service =", err)
			return
		}
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User UnFollowed successfully!!", make(map[string]string))
	}
}
