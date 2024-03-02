package routes

import "github.com/gin-gonic/gin"

func HelloRoute(route *gin.Engine) {
	hello := route.Group("/greeting")
	{
		hello.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello",
			})
		})
		hello.GET("/goodmorning", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "goodmorning",
			})
		})
	}
}
