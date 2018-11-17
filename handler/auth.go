package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type RegisterForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginForm struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/html/login.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		var login LoginForm
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &login)
		email := login.Email
		password := login.Password
		user := db.l.GetUserByEmailPassword(email, password)
		if user != nil {
			token := db.t.AddToken(user.Id)
			body := make(map[string]string)
			body["token"] = string(token)
			ReturnAPIResponse(w, r, 200, "User LoggedIn Successfully!!", body)
			return
		}

		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "Incorrect User credentials!!", make(map[string]string))

	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var register RegisterForm
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(body), &register)
		firstName := register.FirstName
		lastName := register.LastName
		email := register.Email
		password := register.Password
		db.l.Add(firstName, lastName, email, password)
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 200, "User Registered Successfully!!", make(map[string]string))

	}

}
