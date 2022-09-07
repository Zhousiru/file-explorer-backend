package main

import (
	"github.com/Zhousiru/file-explorer-backend/internal/api"
	"github.com/Zhousiru/file-explorer-backend/internal/config"
	"github.com/Zhousiru/file-explorer-backend/internal/log"
)

func main() {
	addr := config.Get(config.K_API_ADDR)
	log.Info("listen and serve on %s", addr)

	err := api.StartServer(addr)
	if err != nil {
		log.Err(err.Error())
	}
}
