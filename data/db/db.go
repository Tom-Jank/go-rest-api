package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

func Setup() {
	config := loadDbConf()

	fmt.Println(config)
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
