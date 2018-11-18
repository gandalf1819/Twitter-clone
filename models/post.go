package models

type Post struct {
	Id     int
	UserId int
	Text   string
}

type UserPosts []Post

func NewUserPosts() UserPosts {
	return make(UserPosts, 0)
}

func (up *UserPosts) AddPost(userId int, text string) Post {
	post := Post{
		Id:     IncrementPostId(*up),
		UserId: userId,
		Text:   text,
	}
	*up = append(*up, post)

	return post
}

func IncrementPostId(up UserPosts) int {
	return len(up) + 1
}
