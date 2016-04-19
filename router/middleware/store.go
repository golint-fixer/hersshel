package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hersshel/hersshel/config"
	"github.com/hersshel/hersshel/store"
	"github.com/hersshel/hersshel/store/datastore"
)

// Store is a middleware function that initializes the Datastore and attaches to
// the context of every http.Request.
func Store(cfg *config.PostgreSQL) gin.HandlerFunc {
	db := datastore.New(cfg)

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
