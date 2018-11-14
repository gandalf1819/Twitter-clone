package main

import (
	"../handler"
	"log"
	"net/http"
)

func main() {

	log.Println("Attempting to connect to server on port 9090")
	handler.Init()
	http.Handle("/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/register/", handler.Register)
	http.HandleFunc("/posts/", handler.Posts)
	http.HandleFunc("/follow/", handler.FollowUser)
	http.HandleFunc("/unfollow/", handler.UnfollowUser)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Failed to connect to server:", err)
	}

}
