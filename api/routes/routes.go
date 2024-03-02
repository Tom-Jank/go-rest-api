package routes

import "github.com/gin-gonic/gin"

func HelloRoute(route *gin.Engine) {
	hello := route.Group("/greeting")
	{
		hello.GET("/hello", helloResponseHandler)
		hello.GET("/goodmorning", goodmorningResponseHandler)
	}
}

func helloResponseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func goodmorningResponseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "goodmorning",
	})
}
