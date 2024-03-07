package routes

import (
	"net/http"

	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/Tom-Jank/go-rest-api/data/models"
	"github.com/gin-gonic/gin"
)

func MonkeRoute(route *gin.Engine) {
	monke := route.Group("/monke")
	{
		monke.GET("", findAllResponseHandler)
		monke.POST("", createNewResponseHandler)
	}
}

func findAllResponseHandler(c *gin.Context) {
	monkes := db.GetAllMonke()

	c.IndentedJSON(http.StatusOK, monkes)
}

func createNewResponseHandler(c *gin.Context) {
	var newMonke models.Monke

	if err := c.BindJSON(&newMonke); err != nil {
		return
	}

}
