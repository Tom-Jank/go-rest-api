package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string
}

func FindAllUsersResponseHandler(c *gin.Context) {
    result,err := findAllUsers(db.DBCon)
    if err != nil {
        log.Panic("Error fetching users", err)
    }

    c.IndentedJSON(http.StatusOK, result)
}

func AddUserRequestHandler(c *gin.Context) {
    var user User
    err := c.Bind(&user)
    if err != nil {
        fmt.Print("nah")
    }
    result := addUser(&user, db.DBCon) 
    handleCreationErr(result, c)
    
}

func findAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil 
}

func addUser(user *User, db *gorm.DB) *gorm.DB {
    return db.Create(user)
}

func handleCreationErr(result *gorm.DB, c *gin.Context) {
    if result.Error != nil {
        switch result.Error {
        case gorm.ErrDuplicatedKey:
            c.IndentedJSON(http.StatusBadRequest, "User with same Id already exists")
        }
    } else {
        c.IndentedJSON(http.StatusOK, "created")
    }
}

