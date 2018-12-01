package main

import (
	ud "../user_driver"
	"../userpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("server started")
	lis, err := net.Listen("tcp", "0.0.0.0:5000")
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
