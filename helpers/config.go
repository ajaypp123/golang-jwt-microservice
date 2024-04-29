package helpers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ajaypp123/golang-jwt-microservice/models"
)

func GetConfig() *models.Config {
	configFile, err := os.Open("conf/config.json")
	if err != nil {
		log.Fatal("Error opening config file:", err)
		return nil
	}
	defer configFile.Close()

	var config models.Config
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatal("Error decoding config JSON:", err)
		return nil
	}
	return &config
}
