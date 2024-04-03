package util

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"math/rand"
	"os"
	"time"
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

func GenerateSnowflake(t time.Time) int64 {
	epoch := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	now := t.UTC().UnixNano() / 1000000
	return (now - epoch) << 22
}

func ParseSnowflake(snowflake int64) time.Time {
	epoch := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	return time.Unix(0, ((snowflake>>22)+epoch)*1000000)
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}
