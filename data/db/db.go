package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
    DBCon *gorm.DB
)

func Setup() {
    var err error
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSL"))
    DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
    catchConnectionError(err)
    log.Print("Connected to db")
}

func catchConnectionError(err error) {
    if err != nil {
        log.Panic("Error connenting to db", err)
    }
}


