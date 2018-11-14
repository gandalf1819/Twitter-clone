package handler

import (
	"../models"
	"html/template"
	"net/http"
)

type PostsPageData struct {
	Friends []models.UserList
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

		//used to post data
	}

}
