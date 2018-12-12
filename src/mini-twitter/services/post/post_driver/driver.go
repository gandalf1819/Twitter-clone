package post_driver

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"mini-twitter/services/post/postpb"
	"mini-twitter/util"
)

type Server struct{}

var up postpb.UserPosts

func Init() {
	log.Println("Init called ")
	_, err := util.InteractWithRaftStorage("PUT", "postDB", up)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
}

func GetPostDB(value interface{}) (postpb.UserPosts, error) {
	var db postpb.UserPosts
	data, err := util.InteractWithRaftStorage("GET", "postDB", db)
	if err != nil {
		log.Println("Error occured while getting post data from Raft =", err)
		panic(err)
	}
	var postDB postpb.UserPosts
	postDB, err = DecodeRaftPostStorage(data)
	if err != nil {
		log.Println("Error occured while decoding post data from Raft storage =", err)
		return postDB, err
	}
	log.Println("postDB after decode =", postDB)
	return postDB, nil
}

func DecodeRaftPostStorage(db string) (postpb.UserPosts, error) {
	log.Println("Decode post Storage called")
	dec := gob.NewDecoder(bytes.NewBufferString(db))
	if err := dec.Decode(&up); err != nil {
		log.Fatalf("raftexample: could not decode message (%v)", err)
		return up, err
	}
	log.Println("postDB in DecodeRaftpostStorage =", up)

	return up, nil
}

func (*Server) AddPost(ctx context.Context, postDetails *postpb.PostText) (*postpb.Post, error) {
	log.Println("AddPost API called")
	var up postpb.UserPosts
	post := &postpb.Post{
		Id:     IncrementPostId(),
		UserId: postDetails.UserId,
		Text:   postDetails.Text,
	}
	postDB, err := GetPostDB(up)
	if err != nil {
		return nil, err
	}
	postDB.Posts = append(postDB.Posts, post)
	log.Println("PostsDB = ", postDB)

	_, err = util.InteractWithRaftStorage("PUT", "postDB", postDB)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
	return post, nil
}

func (*Server) GetFollowerPosts(ctx context.Context, users *postpb.Users) (*postpb.UserPosts, error) {
	log.Println("GetFollowerPosts called")
	var up postpb.UserPosts
	posts := &postpb.UserPosts{
		Posts: make([]*postpb.Post, 0),
	}
	postDB, err := GetPostDB(up)
	if err != nil {
		return nil, err
	}

	for _, user := range users.Ids {
		for _, userPostsObj := range postDB.Posts {
			if user == userPostsObj.UserId {
				posts.Posts = append(posts.Posts, userPostsObj)
			}
		}
	}

	_, err = util.InteractWithRaftStorage("PUT", "postDB", postDB)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}

	return posts, nil
}

func (*Server) GetAllPosts(ctx context.Context, in *postpb.NoArgs) (*postpb.UserPosts, error) {
	var err error
	log.Println("GetAllPosts called")
	up, err = GetPostDB(up)
	if err != nil {
		return nil, err
	}
	return &up, nil
}

func IncrementPostId() int32 {
	return int32(len(up.Posts) + 1)
}
