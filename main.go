package main

import (
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

func main() {
	echo := echo.New()

	echo.GET("/", helloWorld)
	echo.Static("/public", "public")
	// echo.Logger.Fatal(echo.Start(":1323"))

	http.Handle("/public", echo)
	http.Handle("/", echo)
	appengine.Main()
}

func helloWorld(c echo.Context) error {
	type Link struct {
		URL  string
		Type string
	}
	var link []Link
	link = append(link, Link{URL: "https://www.facebook.com/gotgatgat", Type: "Facebook"})
	link = append(link, Link{URL: "https://www.medium.com/@wittayasutt", Type: "Medium"})
	return c.JSON(http.StatusOK, link)
}
