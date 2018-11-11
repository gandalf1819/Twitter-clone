package handler

import (
	"html/template"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/html/login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
		
		
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func Register(w http.ResponseWriter, r *http.Request){
	db.l.Add("Nikhil", "Nar", "ncn251@nyu.edu", "Rubique@1993")

	fmt.Println("Login values====", db.l)
}