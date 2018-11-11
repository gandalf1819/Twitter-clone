package handler

import (
	"log"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/html/login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
		
		//email:= r.Form["email"][0]
		//password:= r.Form["password"][0]

		//db.t.GetToken()
		//Cookie part to be written
    }
}

func Register(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
		r.ParseForm()
		firstName:=r.Form["first_name"][0]
		lastName:=r.Form["last_name"][0]
		email:=r.Form["email"][0]
		password:=r.Form["password"][0]
		db.l.Add(firstName, lastName, email, password)
		log.Println("Login values=======",db.l)
	}
	
}