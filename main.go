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
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))

	e.GET("/", routeHome)

	e.POST("/:index/:type/_doc/:id", routeIndex)
	e.POST("/:index/:type/_docs", routeBatchIndex)

	e.GET("/:index/:type/_doc/:id", routeGet)

	e.POST("/:index/:type/_aggregate/:field/:func", routeAggregate)
	e.POST("/:index/:type/_search", routeSearch)

	e.DELETE("/:index/:type/_doc/:id", routeDelete)

	log.Fatal(e.Start(*flagListenAddr))
}
