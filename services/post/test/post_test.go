package post_test

import (
	"../postpb"
	"context"
	"log"
	"testing"
)

func TestAddPost(t *testing.T) {
	log.Println("Executing TestAddPost TEST CASE")
	InitializePostClient()

	// Add Sample Posts
	var postData []postpb.PostText
	postData = []postpb.PostText {

		postpb.PostText {
			UserId: 1,
			Text: "Nikhil",
		},
		postpb.PostText {
			UserId: 2,
			Text: "Yuvraj",
		},
		postpb.PostText {
			UserId: 3,
			Text: "Chinmay",
		},
	}

	/*for id, value := range postData {

		_, err := pc.PostDB.Add(context.Background(), &value)
		if err != nil {
			t.Errorf("Data with first name %v and last name %v not inserted", value.FirstName, value.LastName)
			log.Println("Error received from User Service =", err)
			return
		}
	}*/
	log.Println("TestRegister Test Passed")

	pc.PostDB.AddPost(1,11,"Post 1")
	pc.PostDB.AddPost(2, 22, "Post 2")
	pc.PostDB.AddPost(3,33,"Post 3")
}

func TestGetFollowerPosts(t *testing.T) {
	log.Println("Executing TestGetFollowerPosts TEST CASE")

	followerMap := make(map[string]int)
	//Nikhil Follows Yuvraj

	fp := &postpb.Post {
		Id: int32(1),
		UserId: int32(2),
		Text: "Nikhil Follows Yuvraj",
	}
	_, err := pc.postDB.PostText(context.Background(), fp)
	if err != nil {
		log.Println("Error received from User Service =", err)
		return
	}

	followerMap["Yuvraj"] = 0
	//Nikhil Follows Chinmay
	fp = &postpb.Post {
		Id: int32(1),
		UserId: int32(3),
		Text: "Nikhil Follows Chinmay",
	}
	_, err = pc.postDB.FollowUser(context.Background(), fp)
	if err != nil {
		log.Println("Error received from User Service =", err)
		return
	}
	followerMap["Chinmay"] = 0

	user := &pb.UserId{
		Id: int32(1),
	}

	log.Println("TestFollowUser Test Passed")
}