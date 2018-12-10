package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	ud "mini-twitter/services/user/user_driver"
	"mini-twitter/services/user/userpb"
	"net"
	"os"
)

func main() {
	fmt.Println("server started")
	userPort := os.Getenv("USER_PORT")
	log.Println("userPort =", userPort)
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
