package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
    DBCon *gorm.DB
)

func Setup() {
    var err error
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    catchError(err)
    log.Print("Connected to db")
}

func catchError(err error) {
    if err != nil {
        log.Panic("Error connenting to db", err)
    }
}


