package main

import (
	"net/http"

	"github.com/Tom-Jank/go-rest-api/api/routes"
	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env")
    db.Setup()
    router := gin.Default()
    routes.UserRoutes(router)
    http.ListenAndServe(":8080", router)
}
