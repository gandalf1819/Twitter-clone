## Webber
---

Webber is a simple web application, comprised of a web server written in Go. 

---

## Features

Webber provides the following features in the application.

1. User registration with username and password.
2. Logging in as a given user, given username and password.
3. Users can follow or unfollow other users.
4. Users can create posts that are associated with their profile.
5. Users can view posts of the users who they follow on their news feed.

---

## Project Schema

Webber 

    |-- handler
    	|-- auth.go --> authentication module, allows user registration and login portal.
    	|-- init.go --> initialization module, initializes the connectors module.
    	|-- post.go --> config for post module, implementation of post, follow and unfollow functionality.
    	|-- response.go --> returns API response to user.
		|-- connectors.go --> connectors module, handles connections to authentication, user and userpost services.
    |-- services
		|-- auth
			|-- auth_driver --> handles token initialization, generation and deletion.
				|-- driver.go
			|-- auth_server --> authentication server module.
				|-- server.go
			|-- auth_test --> authetication testing module, contains test cases for the functions of driver.go.
				|-- client.go
				|-- auth_test.go
                |-- runTestCases.sh --> shell script to start test client and run test cases
			|-- authpb --> gRPC function calls using protocol buffers.
				|-- auth.pb.go
				|-- auth.proto
            |-- runAuthService.sh --> shell script to run Auth Service   
		|-- post
			|-- post_driver --> add new posts and retrieve follower posts functionalities. 
				|-- driver.go
			|-- post_server --> post server module.
				|-- server.go
			|-- test --> post testing module, contains test cases for the functions of driver.go.
				|-- client.go
				|-- post_test.go
                |-- runTestCases.sh --> shell script to start test client and run test cases	
			|-- postpb --> gRPC function calls using protocol buffers.
				|-- post.pb.go
				|-- post.proto
            |-- runPostService.sh --> shell script to run Post Service   
		|-- user
			|-- user_driver --> implementation of user module functionalities.
				|-- driver.go
			|-- user_server --> user server module.
				|-- server.go
			|-- test --> user testing module, contains test cases for the user module functionalities of user.go.
				|-- client.go
				|-- user_test.go
                |-- runTestCases.sh --> shell script to start test client and run test cases
			|-- userpb --> gRPC function calls using protocol buffers.
				|-- user.pb.go
				|-- user.proto
            |-- runUserService.sh --> shell script to run User Service
	|-- raft
		|-- raftexample-1
		|-- raftexample-2
		|-- raftexample-3
		|-- app.exe
		|-- listener.go
		|-- main.go
		|-- raft.go
		|-- raftexample_test.go
	|--views
		|--css
			|-- main.css --> stylesheet for mini-twitter
			|-- toastr.css --> stylesheet for toastr functionality
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
	|-- runBackendServer.sh --> port specifications and command to start the server

---

### RAFT Implementation

Raft is a distributed consensus algorithm. It solves the problem of getting multiple servers to agree on a shared state even in the face of failures. The shared status is usually a data structure supported by a replicated log. We need the system to be fully operational as long as a majority of the servers are up.

Raft works by electing a leader in the cluster. The leader is responsible for accepting client requests and managing the replication of the log to other servers. The data flows only in one direction: from leader to other servers.

Raft decomposes consensus into three sub-problems:

* Leader Election: 
	A new leader needs to be elected in case of the failure of an existing one.

* Log replication: 
	The leader needs to keep the logs of all servers in sync with its own through replication.

* Safety: 
	If one of the servers has committed a log entry at a particular index, no other server can apply a different log entry for that index.

---

## Running the application

Go to the Webber folder and in the terminal run the following command:

./runBackendServer.sh

The PORT property can be changed in the runBackendServer.sh file.

Open you web browser and go to the following address:

https://localhost:9090/login

Use the same port that is defined in the runBackendServer.sh file.

Launch the post, user and auth services with the help of shell scripts file defined in the respective services folder.

To start the raft cluster, go to raft folder and run goreman start.

---


## Team
* Nikhil Nar (ncn251)
* Chinmay Wyawahare (cnw282)
