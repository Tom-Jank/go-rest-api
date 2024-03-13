package routes

import (
	"github.com/Tom-Jank/go-rest-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
    user := route.Group("/user")
    {
        user.GET("", handlers.FindAll)
        user.GET("/:id", handlers.FindByID)
        user.POST("", handlers.AddNew)
    }
}
