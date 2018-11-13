package handler

import (
	"io/ioutil"
	"log"
	"html/template"
	"net/http"
	"encoding/json"
)

type RegisterForm struct{
	FirstName string
	LastName string
	Email string
	Password string
}

type ResponseMessage struct{
	Status int
	Message string
}

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
		var register RegisterForm
		body, err := ioutil.ReadAll(r.Body)
		
		if err !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &register)
		firstName:=register.FirstName
		lastName:=register.LastName
		email:=register.Email
		password:=register.Password
		db.l.Add(firstName, lastName, email, password)

		ReturnAPIResponse(w, r, 200, "User Registered Successfully!!")
		
	}
	
}

func Posts(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET" {
        t, _ := template.ParseFiles("./views/html/posts.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
		
		//email:= r.Form["email"][0]
		//password:= r.Form["password"][0]

		//db.t.GetToken()
		//Cookie part to be written
    }
	
}

func ReturnAPIResponse(w http.ResponseWriter, r *http.Request, status int, message string){
	response:=ResponseMessage{
		Status: status,
		Message: message,
	}
	
	res,err:= json.Marshal(response)
	if err !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("res====",res)
	w.Header().Set("content-type","application/json")
	w.Write(res)
}