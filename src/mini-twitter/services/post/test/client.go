package test

import (
	"google.golang.org/grpc"
	"log"
	"mini-twitter/services/post/postpb"
	"os"
)

type PostClient struct {
	PostDB   postpb.PostServiceClient
	Post     *grpc.ClientConn
	PostPort string
}

var pc PostClient

func InitializePostClient() {
	var err error

	pc.PostPort = os.Getenv("USER_POST_PORT")
	log.Println("postPort =", pc.PostPort)
	pc.Post, err = grpc.Dial("localhost:"+pc.PostPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	pc.PostDB = postpb.NewPostServiceClient(pc.Post)

	log.Println("Post Client created at port =", pc.PostPort)
}
