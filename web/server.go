package main

import (
	"log"
	"net/http"
	"../handler"
)


func main(){

	log.Println("Attempting to connect to server on port 9090")
	handler.Init()
	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/register/", handler.Register)
	
	
	
	if err:= http.ListenAndServe(":9090",nil); err!=nil{
		log.Fatal("Failed to connect to server:", err)
	}
	
	

}