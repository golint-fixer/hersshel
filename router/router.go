package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hersshel/hersshel/api"
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
		v1.POST("/feeds", api.PostFeed)
		v1.POST("/categories", api.PostCategory)
		v1.POST("/items", api.UpdateFeeds)

		v1.GET("/items", api.GetAllItems)
		v1.GET("/feeds", api.GetAllFeeds)
		v1.GET("/feeds/:id/items", api.GetFeedItems)
		v1.GET("/categories/:id/items", api.GetCategoryItems)

		v1.PATCH("/feeds")
		v1.PATCH("/categories")

		v1.DELETE("/feeds/:id", api.DeleteFeed)
	}
	return e
}
