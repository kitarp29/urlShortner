package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lithammer/shortuuid/v3"
)

var ListMap = map[string](string){"kitarp29": "https://twitter.com/kitarp29"}

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, Redirect(c))
	})

	e.POST("addlink", func(c echo.Context) error {
		return c.String(http.StatusOK, AddLink(c))
	})

	e.GET("list", func(c echo.Context) error {
		return c.String(http.StatusOK, List(c))
	})

	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}

func Redirect(c echo.Context) string {
	var response string
	log.Println(c.QueryParam("u"))
	uuid := c.QueryParam("u")
	for key, value := range ListMap {
		if key == uuid {
			return value
		}
	}
	response = "No URL found"
	return response
}

func AddLink(c echo.Context) string {
	var response string
	link := c.QueryParam("link")
	newkey := shortuuid.New()
	ListMap[newkey] = link
	response += "http://localhost:8000/?u=" + newkey
	return response
}

func List(c echo.Context) string {
	var response string
	for key, value := range ListMap {
		response = response + key + " : " + value
	}
	return response
}
