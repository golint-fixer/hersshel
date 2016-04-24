package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hersshel/hersshel/engine"
)

// Engine is a middleware function that initializes
// an Engine and attach it to the context.
func Engine() gin.HandlerFunc {
	var once sync.Once
	var e engine.Engine

	return func(c *gin.Context) {
		once.Do(func() {
			e = engine.New()
		})

		engine.ToContext(c, e)
		c.Next()
	}
}
