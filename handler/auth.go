package handler

import (
	"../services/auth/authpb"
	"../services/user/userpb"
	"context"
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
		//user := db.l.GetUserByEmailPassword(email, password)
		loginDetails := &userpb.LoginDetails{
			Email:    email,
			Password: password,
		}
		user, err := con.GetUserClient().GetUserByEmailPassword(context.Background(), loginDetails)
		if err != nil {
			ReturnAPIResponse(w, r, 422, "Error occured while login. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from User Service =", err)
			return
		}
		if user.Id != 0 {
			//token := db.t.AddToken(user.Id)
			userId := &authpb.UserId{
				User: int32(user.Id),
			}
			body := make(map[string]string)
			token, err := con.GetAuthTokenClient().AddToken(context.Background(), userId)
			if err != nil {
				ReturnAPIResponse(w, r, 422, "Error occured while login. Contact your system admin for more details!!", body)
				log.Println("Error received from Auth Service =", err)
				return
			}

			body["Token"] = token.TokenName
			log.Println("db.t=======", db.t)
			log.Println("db.l===", db.l)

			tokCook := &http.Cookie{
				Name:    "token",
				Value:   token.TokenName,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			}

			userCookie := &http.Cookie{
				Name:    "user_id",
				Value:   strconv.Itoa(int(user.Id)),
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
		userParams := &userpb.AddUserParameters{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
		}
		_, err = con.GetUserClient().Add(context.Background(), userParams)
		if err != nil {
			ReturnAPIResponse(w, r, 422, "Error occured while login. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from User Service =", err)
			return
		}
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

		tokName := &authpb.AuthTokenName{
			TokenName: cookieToken.Value,
		}

		status, err := con.GetAuthTokenClient().UnsetToken(context.Background(), tokName)
		if err != nil || !status.ResponseStatus {
			ReturnAPIResponse(w, r, 422, "Error occured while logout. Contact your system admin for more details!!", make(map[string]string))
			log.Println("Error received from Auth Service =", err)
			return
		}
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
