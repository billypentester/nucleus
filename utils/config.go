package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"nucleus/models"
	"os"
)

func LoadConfig() (*models.Config, error) {

	configPath := os.Getenv("CONFIG_PATH")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	defer file.Close()

	log.Println("- Config file loaded from: ", configPath)

	var config models.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	return &config, nil
}
