package util

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Dsn    string `toml:"dsn"`
		Driver string `toml:"driver"`
	} `toml:"database"`
	Server struct {
		IP   string `toml:"ip"`
		Port int    `toml:"port"`
	} `toml:"server"`
}

func LoadConfigFile() (config Config) {
	file, err := os.ReadFile("../../config.toml")
	if err != nil {
		log.Fatalf("Error reading TOML file: %v", err)
	}
	if err = toml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Error unmarshaling TOML: %v", err)
	}
	return config
}
