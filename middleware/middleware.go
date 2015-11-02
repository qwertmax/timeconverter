// Copyright © 2014, 2015 Good Dog Labs., Inc.
// All Rights Reserved.

package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middleware witch allow foreing requests.
//
// for CORS we must accept request type OPTIONS and return status code = 200
//
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: expand other parameters as config values
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
