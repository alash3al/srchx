package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/", routeHome)
	e.POST("/:index/:type/_doc/:id", routeIndex)
	e.POST("/:index/:type/_docs", routeBatchIndex)
	e.GET("/:index/:type/_doc/:id", routeGet)
	e.DELETE("/:index/:type/_doc/:id", routeDelete)
	e.Match([]string{"GET", "POST"}, "/:index/:type/_search", routeSearch)

	log.Fatal(e.Start(*flagListenAddr))
}
