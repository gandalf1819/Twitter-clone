package main

import (
	pd "../post_driver"
	"../postpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("server started")
	postPort := os.Getenv("USER_POST_PORT")
	log.Println("postPort =", postPort)
	lis, err := net.Listen("tcp", "0.0.0.0:"+postPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &pd.Server{})

	pd.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
