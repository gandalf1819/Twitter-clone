package handler

import (
	"../models"
)

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

	InitializeConnectors()
}
