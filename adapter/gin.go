package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinAdapter adapts an http.Handler to a Gin handler.
func GinAdapter(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// GinAdapterFunc adapts an http.HandlerFunc to a Gin handler.
func GinAdapterFunc(handlerFunc http.HandlerFunc) gin.HandlerFunc {
	return GinAdapter(handlerFunc)
}
