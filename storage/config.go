package storage

import (
	"encoding/json"
	"os"
)

type GodoConfig struct {
	FilePath string
}

func readConfig() GodoConfig {
	config := GodoConfig{
		FilePath: "file_storage",
	}
	configPath, ok := os.LookupEnv("GODO_CONFIG")
	if !ok {
		configPath = "~/.godoconfig"
	}
	file, _ := os.OpenFile(configPath, os.O_CREATE|os.O_RDONLY, 0644)
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
	return config
}
