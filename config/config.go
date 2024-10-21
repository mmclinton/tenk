package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ApiKey string `json:"api_key"`
}

func ReadConfig() (*Config, error) {
	configFile := filepath.Join(os.Getenv("HOME"), ".config", "tenk", "config.json")

	f, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
