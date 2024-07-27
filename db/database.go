package db

import (
	"database/sql"
	"gosurpher/models"

	_ "github.com/mattn/go-sqlite3"
)

var db_type string = "sqlite3"
var db_path string = "./db/gosurpher.db"

func Select_blogs() []models.Blog {

	db, _ := sql.Open(db_type, db_path)

	// query
	rows, _ := db.Query("SELECT * FROM Blog")

	var blogs []models.Blog

	for rows.Next() {
		var blog models.Blog

		rows.Scan(&blog.Id, &blog.Date, &blog.Title, &blog.Autor, &blog.Content, &blog.Route, &blog.Type)

		blogs = append(blogs, blog)
	}

	rows.Close()
	db.Close()

	return (blogs)

}

func Select_blog_by_route(route string) models.Blog {

	db, _ := sql.Open(db_type, db_path)

	rows, _ := db.Query("SELECT * FROM Blog WHERE Route = \"" + route + "\"")
	rows.Next()

	var blog models.Blog
	rows.Scan(&blog.Id, &blog.Date, &blog.Title, &blog.Autor, &blog.Content, &blog.Route, &blog.Type)

	rows.Close()
	db.Close()

	return (blog)

}
