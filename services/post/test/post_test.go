package test

import (
	"../postpb"
	"context"
	"log"
	"testing"
)

func TestAddPost(t *testing.T) {
	log.Println("Executing TestAddPost TEST CASE")
	InitializePostClient()

	postsMap := make(map[string]int)
	//Yuvraj adds status
	userPost := &postpb.PostText{
		UserId: int32(3),
		Text:   "This is a Yuvraj's status",
	}
	_, err := pc.PostDB.AddPost(context.Background(), userPost)
	if err != nil {
		log.Println("Error received from UserPost Service =", err)
		return
	}
	postsMap["Yuvraj"] = 0

	//Nikhil adds status
	userPost = &postpb.PostText{
		UserId: int32(1),
		Text:   "This is a Nikhil's status",
	}
	_, err = pc.PostDB.AddPost(context.Background(), userPost)
	if err != nil {
		log.Println("Error received from UserPost Service =", err)
		return
	}
	postsMap["Nikhil"] = 0

	//Chinmay adds status
	userPost = &postpb.PostText{
		UserId: int32(2),
		Text:   "This is a Chinmay's status",
	}
	_, err = pc.PostDB.AddPost(context.Background(), userPost)
	if err != nil {
		log.Println("Error received from UserPost Service =", err)
		return
	}
	postsMap["Chinmay"] = 0

	var allPosts *postpb.UserPosts
	allPosts, err = pc.PostDB.GetAllPosts(context.Background(), &postpb.NoArgs{})
	if err != nil {
		log.Println("Error received from UserPost Service =", err)
		return
	}

	for _, value := range allPosts.Posts {
		if value.UserId == 1 {
			postsMap["Nikhil"] = 1
		} else if value.UserId == 2 {
			postsMap["Chinmay"] = 1
		} else if value.UserId == 3 {
			postsMap["Yuvraj"] = 1
		}
	}

	for key, value := range postsMap {
		if value != 1 {
			t.Errorf("%v failed to add post", key)
		}

	}
	log.Println("TestRegister Test Passed")
}
