package models

type Post struct{
	id int
	userId int
	text string
}

type UserPosts []Post

func NewUserPosts()UserPosts{
	return make(UserPosts,0)
}

func (up *UserPosts)AddPost(userId int,text string) Post{
	post:=Post{
		id: IncrementPostId(*up),
		userId: userId,
		text: text,
	}
	*up = append(*up, post)

	return post
}

func IncrementPostId(up UserPosts) int{
	return len(up) + 1
}