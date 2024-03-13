package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tom-Jank/go-rest-api/data/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string
}

func FindAll(c *gin.Context) {
    result,err := findAllUsers(db.DBCon)
    if err != nil {
        log.Panic("Error fetching users", err)
    }
    c.IndentedJSON(http.StatusOK, result)
}

func AddNew(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := addUser(&user, db.DBCon); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.IndentedJSON(http.StatusOK, user)
}

func FindByID(c *gin.Context) {
    ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Provided bad ID"})
    }
    user, err := findUserByID(uint(ID), db.DBCon)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Provided bad ID"})
    } 
    c.IndentedJSON(http.StatusOK, gin.H{"success": user})
}

func findAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil 
}

func findUserByID(ID uint, db *gorm.DB) (User, error) {
    var user User
    if err := db.First(&user, ID).Error; err != nil {
        return user, err
    }
    return user, nil
}

func addUser(user *User, db *gorm.DB) error {
    result := db.Create(user)
    return result.Error
}

