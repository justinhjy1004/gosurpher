package main

import (
	"context"
	"gosurpher/handlers"
	"log"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Static CSS
	e.Static("/", "static/")

	e.GET("/", handlers.HomeHandler)
	e.GET("/blog", handlers.BlogHandler)
	e.Logger.Fatal(e.Start(":42069"))

	// Use the application default credentials
	projectID := "GoSurpher"
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

}
