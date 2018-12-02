package test

import (
	"../postpb"
	"google.golang.org/grpc"
	"log"
)

type PostClient struct {
	PostDB postpb.PostServiceClient
	Post *grpc.ClientConn
	PostPort string
}

var pc PostClient

func InitializePostClient() {
	var err error

	pc.PostPort = "5000"
	pc.Post, err = grpc.Dial("localhost:"+pc.PostPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	pc.PostDB = postpb.NewPostServiceClient(pc.Post)

	log.Println("Post Client created at port =", pc.PostPort)
}

