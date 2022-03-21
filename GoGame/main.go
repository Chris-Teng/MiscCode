package main

import (
	"game/api"

	"github.com/gin-gonic/gin"
)

func main() {
	//router server
	r := gin.Default()
	r.GET("/", CORSMiddleware(), func(c *gin.Context) {
		c.JSON(200, `{"msg":"Hello,Go!"}`)
	})
	// r.POST("/login", api.Login)

	r.GET("/queryStatistics", api.QueryStatistics)
	r.POST("/attack", api.Attack)
	r.Run(":8000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
