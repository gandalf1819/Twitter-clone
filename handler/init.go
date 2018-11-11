package handler

import(
	"../models"
)

type database struct{
	l models.Login
}

var db database

func Init(){
	l:= models.NewLogin()
	db = database{l}
}
