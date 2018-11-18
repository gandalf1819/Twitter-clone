## Mini Twitter
---

## About

Mini twitter is a simple web application, comprised of a web server written in Go. For stage 1 of this project, instead of using a database, we are keeping everything in memory. 

---

## Features

Mini-twitter provides the following features in the application.

1. User registration with username and password.
2. Logging in as a given user, given username and password.
3. Users can follow or unfollow other users.
4. Users can create posts that are associated with their profile.
5. Users can view posts of the users who they follow on their news feed.

---

## Project Schema

mini-twitter 

    |-- handler
    	|-- auth.go --> authentication module, allows user registration and login portal.
    	|-- init.go --> initialization module, initializes the login, token and post models.
    	|-- post.go --> config for post module, implementation of post, follow and unfollow functionality.
    	|-- response.go --> returns API response to user.
    |-- models
        |-- post.go   --> post model, defines post module features.
        |-- token.go --> token validity duration, hash difficulties, etc.
        |-- user.go --> user model, defines user module features.
	|--views
		|--css
			|-- main.css --> stylesheet for mini-twitter
			|-- toastr.css --> stylesheet for toastr functionality
			|-- toastr.js.map
			|-- toastr.min.css
			|-- toastr.min.js
		|--html
			|-- login.html --> homepage containing user login and user registration features
			|-- posts.html --> posts page containing friends, newsfeed and posts features
		|--images
		|--js
			|-- api.js --> function definitions for user registration, signIn, followUser, unfollowUser, addPost, logout
			|-- login.js --> function definition for toggleTab for loginForm, registerForm
			|-- posts.js --> function definition for toggleTab for friends, newsfeed and status tabs
			|-- toastr.js --> function definitions for toastr features
	|--web
		|--server.go --> build target (the code that actually runs the server)

---

## Running the application

Go to the mini-twitter folder and in the terminal run the following command:

./runBackendServer.sh

Open you web browser and go to the following address:

https://localhost:9090/login

---


## Team
* Nikhil Nar (ncn251)
* Chinmay Wyawahare (cnw282)