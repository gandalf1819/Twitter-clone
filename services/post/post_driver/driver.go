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

func (*Server) GetFollowerPosts(ctx context.Context, users *postpb.Users) (*postpb.UserPosts, error) {
	posts := &postpb.UserPosts{
		Posts: make([]*postpb.Post, 0),
	}

	for _, user := range users.Ids {
		for _, userPostsObj := range up.Posts {
			if user == userPostsObj.UserId {
				posts.Posts = append(posts.Posts, userPostsObj)
			}
		}
	}

	return posts, nil
}

func IncrementPostId() int32 {
	return int32(len(up.Posts) + 1)
}
