package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	Title   string
	Content string
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5454/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	_, err = db.Exec("INSERT INTO posts VALUES($1, $2)", "Title 2", "Content 2")
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("SELECT * FROM posts;")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	for result.Next() {
		var p Post
		err := result.Scan(&p.Title, &p.Content)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", p)
	}
}
