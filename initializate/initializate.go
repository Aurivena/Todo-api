package initializate

import (
	"Todo/models"
	"encoding/json"
	"os"
)

const (
	configFilePath = "config.json"
)

var (
	ConfigService = &models.ConfigService{}
)

func Load() error {
	if err := loadConfiguration(); err != nil {
		return err
	}
	return nil
}

func loadConfiguration() error {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &ConfigService)
	if err != nil {
		return err
	}

	return nil
}
