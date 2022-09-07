package config

import (
	"os"

	"github.com/Zhousiru/file-explorer-backend/internal/log"
	"github.com/buger/jsonparser"
)

const configPath = "./config.json"

var configData []byte

const (
	K_API_ADDR = "apiAddr"
	K_ROOT     = "root"
)

func init() {
	Load()
}

func Load() {
	_, err := os.Stat(configPath)
	if err != nil {
		log.Err("config not found")
		os.Exit(-1)
	}

	configData, err = os.ReadFile(configPath)
	if err != nil {
		log.Err(err.Error())
		os.Exit(-1)
	}
}

func Get(k string) string {
	v, err := jsonparser.GetString(configData, k)
	if err != nil {
		log.Err("can't get config: %s", err.Error())
		os.Exit(-1)
	}

	return v
}
