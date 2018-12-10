package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	ad "mini-twitter/services/auth/auth_driver"
	"mini-twitter/services/auth/authpb"
	"net"
	"os"
)

func main() {
	fmt.Println("server started")
	authPort := os.Getenv("AUTH_PORT")
	lis, err := net.Listen("tcp", "0.0.0.0:"+authPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthTokenServiceServer(s, &ad.Server{})

	ad.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
