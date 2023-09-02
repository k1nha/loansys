package main

import (
	"github.com/lucascmpus/lend/config"
	"github.com/lucascmpus/lend/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//Initialize configs
	err := config.Init()

	if err != nil {
		logger.ErrorF("config initalizaion error: %v", err)
		return
	}

	//Initialize router
	router.InitializeRouter()
}
