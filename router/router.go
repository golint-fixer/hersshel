package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hersshel/hersshel/router/middleware/header"
)

// Load takes infinite number of middleware and apply them in order
// to the gin router.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(header.NoCache)
	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middleware...)

	v1 := e.Group("/v1")
	{
		v1.POST("/feeds")
		v1.POST("/categories")

		v1.GET("/items")
		v1.GET("/feeds")
		v1.GET("/feeds/:feed_id/items")
		v1.GET("/categories/:category_id/items")

		v1.PATCH("/feeds")
		v1.PATCH("/categories")
	}
	return e
}
