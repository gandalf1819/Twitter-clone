package main

import (
	"../handler"
	"log"
	"net/http"
	"os"
)

func main() {

	handler.Init()
	http.Handle("/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/register/", handler.Register)
	http.HandleFunc("/posts/", handler.Posts)
	http.HandleFunc("/follow/", handler.FollowUser)
	http.HandleFunc("/unfollow/", handler.UnfollowUser)
	http.HandleFunc("/logout/", handler.LogoutUser)
	log.Println("Attempting to start server on port :", os.Getenv("PORT"))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("SERVER: Failed to listen and serve on port {%s}.  Error message: {%v}", os.Getenv("PORT"), err.Error())
	}

}
