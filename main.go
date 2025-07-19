package main

import (
	"github.com/ojCezarFerreira/goportunities/config"
	"github.com/ojCezarFerreira/goportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
