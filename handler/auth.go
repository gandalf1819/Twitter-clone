package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
		if user.Id != 0 {
			token := db.t.AddToken(user.Id)
			body := make(map[string]string)
			body["Token"] = string(token)
			log.Println("db.t=======", db.t)
			log.Println("db.l===", db.l)

			tokCook := &http.Cookie{
				Name:    "token",
				Value:   token,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			}

			userCookie := &http.Cookie{
				Name:    "user_id",
				Value:   strconv.Itoa(user.Id),
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			}

			// Set the cookies
			http.SetCookie(w, tokCook)
			http.SetCookie(w, userCookie)

			ReturnAPIResponse(w, r, 200, "User LoggedIn Successfully!!", body)
			return
		}
		log.Println("db.l===", db.l)
		ReturnAPIResponse(w, r, 422, "Incorrect User credentials!!", make(map[string]string))

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

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		cookieToken, err := r.Cookie("token")
		if err != nil || cookieToken.Value == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
			log.Printf("HANDLERS-VALIDATE: Failed & Redirected")
			return
		}

		db.t.UnsetToken(cookieToken.Value)
		log.Println("db.t====", db.t)

		tokCook := &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(-100 * time.Hour),
			MaxAge:  -1,
			Path:    "/",
		}

		userCookie := &http.Cookie{
			Name:    "user_id",
			Value:   "",
			Expires: time.Now().Add(-100 * time.Hour),
			MaxAge:  -1,
			Path:    "/",
		}

		// Set the cookies
		http.SetCookie(w, tokCook)
		http.SetCookie(w, userCookie)
		ReturnAPIResponse(w, r, 200, "User Logged out successfully!!", make(map[string]string))

	}

}
