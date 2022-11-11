package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

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
	log.SetFormatter(&log.JSONFormatter{})
	var log = logrus.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		log.WithFields(logrus.Fields{
			"route":      "/",
			"QureyParam": c.QueryParam("u"),
		}).Info("Urls Shortner called for:" + c.QueryParam("u"))
		if c.QueryParam("u") == "" {
			log.Warn("No Query Param")
			return c.String(http.StatusOK, "NULL Value Called")
		} else {
			response := Redirect(c)
			if response == "No URL found" {
				log.Warn("No URL found")
				return c.String(http.StatusOK, "No URL found")
			} else {
				return c.Redirect(http.StatusMovedPermanently, response)
			}
		}
	})

	e.POST("addlink", func(c echo.Context) error {
		log.WithFields(logrus.Fields{
			"route":      "/addlink",
			"QureyParam": c.QueryParam("link"),
		}).Info("Urls Shortner requested for:" + c.QueryParam("link"))

		if c.QueryParam("link") == "" {
			log.Warn("No Query Param")
			return c.String(http.StatusOK, "NULL Value Called")
		} else {
			return c.String(http.StatusOK, AddLink(c))
		}

	})

	e.GET("list", func(c echo.Context) error {
		log.WithFields(logrus.Fields{
			"route": "/list",
		}).Info("All Short URLs are requested")
		response := List(c)
		if response == "" {
			log.Warn("No Shortened URL found")
			return c.String(http.StatusOK, "No URL Found")
		} else {
			return c.String(http.StatusOK, response)
		}
	})

	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}

func Redirect(c echo.Context) string {
	var response string
	uuid := c.QueryParam("u")

	log.WithFields(logrus.Fields{
		"function": "Redirect",
		"value":    uuid,
	}).Info()

	for key, value := range ListMap {
		if key == uuid {
			return value
		}
	}
	log.WithFields(logrus.Fields{
		"function": "Redirect",
		"value":    uuid,
	}).Warn("No URL Found")
	response = "No URL found"
	return response
}

func AddLink(c echo.Context) string {
	var response string
	link := c.QueryParam("link")
	log.WithFields(logrus.Fields{
		"function": "AddLink",
		"value":    link,
	}).Info()
	newkey := shortuuid.New()
	ListMap[newkey] = link
	response += "http://localhost:8000/?u=" + newkey
	return response
}

func List(c echo.Context) string {
	var response string
	log.WithFields(logrus.Fields{
		"function": "List",
	}).Info()
	for key, value := range ListMap {
		response = response + key + " : " + value + "\n"
	}
	return response
}
