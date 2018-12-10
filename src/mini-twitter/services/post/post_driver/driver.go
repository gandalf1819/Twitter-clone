package post_driver

import (
	"bytes"
	"context"
	"encoding/gob"
	"io/ioutil"
	"log"
	"mini-twitter/services/post/postpb"
	"net/http"
	"strings"
)

type Server struct{}

var up postpb.UserPosts

func Init() {
	log.Println("Init called ")
	_, err := InteractWithRaftStorage("PUT", "postDB", up)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
}

func GetPostDB(value interface{}) (postpb.UserPosts, error) {
	var db postpb.UserPosts
	data, err := InteractWithRaftStorage("GET", "postDB", db)
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

func InteractWithRaftStorage(method string, key string, value interface{}) (string, error) {
	log.Println("Interacted with Raft, method called =", method)
	var payloadValue string
	if method != "GET" {
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(value); err != nil {
			log.Println("Error occured while encoding ", key, " data =", err)
			return "", err
		}
		payloadValue = buf.String()
	}

	url := "http://127.0.0.1:12380/" + key
	var payload *strings.Reader
	payload = nil
	if value != nil {
		payload = strings.NewReader(payloadValue)
	}

	req, _ := http.NewRequest(method, url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error received from Raft =", err)
		return "", err
	}

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error occured while decoding response from Raft =", err)
		return "", err
	}

	log.Println("data received from Raft after calling ", method, " method =", string(data))

	return string(data), nil
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

	_, err = InteractWithRaftStorage("PUT", "postDB", postDB)
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

	_, err = InteractWithRaftStorage("PUT", "postDB", postDB)
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
