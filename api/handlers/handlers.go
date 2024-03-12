package handlers

import (
	"log"
	"net/http"

	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
    ID uint
    Name string
}

func FindAllUsersResponseHandler(c *gin.Context) {
    users,err := findAllUsers(db.DBCon)
    if err != nil {
        log.Panic("Error fetching users", err)
    }

    c.IndentedJSON(http.StatusOK, users)
}

func findAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil 
}
