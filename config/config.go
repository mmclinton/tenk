package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	ApiKey string `json:"api_key"`
}

func ReadConfig() (*Config, error) {
	var configPath string

	if runtime.GOOS == "windows" {
		configPath = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "tenk", "config.json")
	} else {
		configPath = filepath.Join(os.Getenv("HOME"), ".config", "tenk", "config.json")
	}

	f, err := os.Open(configPath)
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
