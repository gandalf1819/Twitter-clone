package main

import (
	ud "../user_driver"
	"../userpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	//"os"
)

func main() {
	fmt.Println("server started")
	userPort := "5000"
	lis, err := net.Listen("tcp", "0.0.0.0:"+userPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &ud.Server{})

	ud.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
