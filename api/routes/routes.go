package routes

import (
	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
)

func HelloRoute(route *gin.Engine) {
	hello := route.Group("/greeting")
	{
		hello.GET("/hello", helloResponseHandler)
		hello.GET("/goodmorning", goodmorningResponseHandler)
	}
}

func MonkeRoute(route *gin.Engine) {
	monke := route.Group("/monke")
	{
		monke.GET("/all", monkeAllResponseHandler)
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

func monkeAllResponseHandler(c *gin.Context) {
	query := db.GetAllMonke()

	c.JSON(200, gin.H{
		"message": query,
	})
}
