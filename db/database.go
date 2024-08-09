package db

import (
	"database/sql"
	"gosurpher/models"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

/* This handles the relevant function calls
for the database for the application */

var db_type string = "sqlite3"
var db_path string = "./db/gosurpher.db"

// Select all blogs and return all the content in them
func Select_blogs() []models.Blog {

	db, _ := sql.Open(db_type, db_path)

	// query
	query := "SELECT * FROM Blog WHERE Id IN (SELECT Id FROM Blog ORDER BY RANDOM() LIMIT 10)"
	rows, _ := db.Query(query)

	var blogs []models.Blog

	for rows.Next() {
		var blog models.Blog

		rows.Scan(&blog.Id, &blog.Date, &blog.Title, &blog.Autor, &blog.Content, &blog.Route, &blog.Type)

		blogs = append(blogs, blog)
	}

	rows.Close()
	db.Close()

	// Randomly render the order of the blogs
	rand.Shuffle(len(blogs), func(i, j int) {
		blogs[i], blogs[j] = blogs[j], blogs[i]
	})

	return (blogs)

}

// Select Blog based on Route
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
