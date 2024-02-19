package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

func Setup() {
	config := loadDbConf()
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config["host"], config["port"], config["user"], config["password"], config["dbname"], config["sslmode"])
	db, err := sql.Open("postgres", conStr)
	errorCatch(err)

	defer db.Close()

	err = db.Ping()
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

func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}
