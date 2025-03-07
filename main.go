package main

import (
	"github.com/costamarcus/go-portunities/config"
	"github.com/costamarcus/go-portunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initilization error: %v", err)
		return
	}

	// initialize router
	router.Initializa()
}
