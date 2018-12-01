package post_driver

import (
	"../postpb"
	"context"
	"log"
)

type Server struct{}

var up postpb.UserPosts

func Init() {
	up := NewUserPosts()

	log.Println("DB Posts Initialized =", up.Posts)
}

func NewUserPosts() *postpb.UserPosts {
	userPosts := &postpb.UserPosts{
		Posts: make([]*postpb.Post, 0),
	}
	log.Println("DB UserPosts Initialized =", up.Posts)
	return userPosts
}

func (*Server) AddPost(ctx context.Context, postDetails *postpb.PostText) (*postpb.Post, error) {
	post := &postpb.Post{
		Id:     IncrementPostId(),
		UserId: postDetails.UserId,
		Text:   postDetails.Text,
	}
	up.Posts = append(up.Posts, post)
	log.Println("PostsDB = ", up.Posts)
	return post, nil
}

func IncrementPostId() int32 {
	return int32(len(up.Posts) + 1)
}
