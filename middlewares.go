package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusNoContent, gin.H{
				"error": http.StatusText(http.StatusNoContent),
			})
			// c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// JSONErrorMiddleware ensures all errors are returned in JSON format
// func JSONErrorMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Next() // Process request

// 		// Check if there are any errors
// 		if len(c.Errors) > 0 {
// 			c.JSON(-1, gin.H{
// 				"error": c.Errors[0].Error(),
// 			})
// 			return
// 		}

// 		// Handle default errors (404, 405, etc.)
// 		statusCode := c.Writer.Status()
// 		if statusCode >= 400 {
// 			c.JSON(statusCode, gin.H{
// 				"error": http.StatusText(statusCode),
// 			})
// 			c.Abort()
// 		}
// 	}
// }
