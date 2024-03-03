package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Tom-Jank/go-rest-api/data/models"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

var database *sql.DB

func Setup() {
	var err error

	config := loadDbConf()
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config["host"], config["port"], config["user"], config["password"], config["dbname"], config["sslmode"])
	database, err = sql.Open("postgres", conStr)
	errorCatch(err)

	err = database.Ping()
	errorCatch(err)

	fmt.Println("Connected to database!")
}

func loadDbConf() map[string]interface{} {
	config := make(map[string]interface{})
	yamlFile, err := os.ReadFile("../db_config.yaml")
	errorCatch(err)

	err = yaml.Unmarshal(yamlFile, config)
	errorCatch(err)

	return config
}

func GetAllMonke() []models.Monke {
	rows, err := database.Query(`select id, monke_name, gender, stand from baboon`)

	if err != nil {
		log.Fatal("DUPA DUPA", err)
	}

	defer rows.Close()

	var monkes []models.Monke

	for rows.Next() {
		var m models.Monke
		err := rows.Scan(&m.ID, &m.Name, &m.Gender, &m.Stand)
		errorCatch(err)
		monkes = append(monkes, m)
	}

	fmt.Println(monkes)
	return monkes
}

func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}
