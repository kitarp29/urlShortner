package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lithammer/shortuuid/v3"
)

var ListMap = map[string](string){}

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
		return Redirect(c)
	})

	e.POST("addlink", func(c echo.Context) error {
		return AddLink(c)
	})

	e.GET("list", func(c echo.Context) error {
		return List(c)
	})

	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}

func Redirect(c echo.Context) error {
	var response string
	uuid := c.QueryParam("u")
	for key, value := range ListMap {
		if key == uuid {
			return c.Redirect(http.StatusMovedPermanently, value)
		} //response = response + key + " : " + value
	}
	response = "No URL found"
	return c.String(http.StatusOK, response)
}

func AddLink(c echo.Context) error {
	var response string
	link := c.QueryParam("link")
	newkey := shortuuid.New()
	ListMap[newkey] = link
	response += "http://localhost:8000/?u=" + newkey
	return c.String(http.StatusOK, response)
}

func List(c echo.Context) error {
	var response string
	for key, value := range ListMap {
		response = response + key + " : " + value
	}
	return c.String(http.StatusOK, response)
}
