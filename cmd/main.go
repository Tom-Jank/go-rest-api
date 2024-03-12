package main

import (
	"net/http"

	"github.com/Tom-Jank/go-rest-api/api/routes"
	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Setup()
	router := gin.Default()
    routes.UserRoutes(router)
    http.ListenAndServe(":8080", router)
}
