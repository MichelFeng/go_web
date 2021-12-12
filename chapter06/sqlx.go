package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp database=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)

	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts(content, author) values($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {

	post := Post{Content: "Hello", AuthorName: "Michel"}
	err := post.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)
	readPost, err := GetPost(post.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(readPost)

}
