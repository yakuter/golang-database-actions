package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Post struct {
	Id    int
	Title  string
	Body string
}

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "godb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db

}

func getAll() {

	db := dbConn()

	selDB, err := db.Query("SELECT * FROM Posts ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	post := Post{}
	posts := []Post{}

	for selDB.Next() {

		var id int
		var title, body string

		err = selDB.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id 	= id
		post.Title 	= title
		post.Body 	= body

		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}
	
	defer db.Close()
}

func getOne(postId int) {

	db 		:= dbConn()

	selDB, err := db.Query("SELECT * FROM Posts WHERE id=?", postId)
	if err != nil {
		panic(err.Error())
	}

	post := Post{}

	for selDB.Next() {

		var id int
		var title, body string

		err = selDB.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id 	= id
		post.Title 	= title
		post.Body 	= body

	}

	fmt.Println("Post Title	: " + post.Title)
	fmt.Println("Post Body	: " + post.Body)

	defer db.Close()
}

func add() {

	db := dbConn()

	title 	:= "Hello Second World"
	body 	:= "The content of the hello second world"
	insertQuery, err := db.Prepare("INSERT INTO Posts(title, body) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(title, body)

	fmt.Println("ADDED: Title: " + title + " | Body: " + body)

	defer db.Close()

}

func update(postId int) {

	db := dbConn()

	title 	:= "Hello 1 World"
	body 	:= "The content of the hello 1 world"
	updateQuery, err := db.Prepare("UPDATE Posts SET title=?, body=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(title, body, postId)

	fmt.Println("UPDATED: Title: " + title + " | Body: " + body)

	defer db.Close()

}

func delete(postId int) {

	db := dbConn()

	deleteQuery, err := db.Prepare("DELETE FROM Posts WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(postId)

	fmt.Println("DELETED")

	defer db.Close()

}

func main()  {

	//add()
	//update(1)
	//delete(1)
	//getOne(1)
	getAll()
	
}