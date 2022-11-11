package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
		var response string

		ListMap["kitarp"] = "www.twitter.com/kitarp"
		for key, value := range ListMap {
			response = response + key + " : " + value
		}
		return c.String(http.StatusOK, response)
	})

	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}
