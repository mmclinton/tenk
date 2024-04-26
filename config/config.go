package config

import (
	"encoding/json"
	"io"
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

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
