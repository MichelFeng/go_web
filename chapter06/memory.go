package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello1", Author: "Author1"}
	post2 := Post{Id: 2, Content: "Hello2", Author: "Author2"}
	post3 := Post{Id: 3, Content: "Hello3", Author: "Author3"}
	post4 := Post{Id: 4, Content: "Hello4", Author: "Author1"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Author1"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Author2"] {
		fmt.Println(post)
	}
}
