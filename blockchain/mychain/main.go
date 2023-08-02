package main

import (
	"fmt"
	"os"

	"github.com/thegreatforge/gokit/config"
	"github.com/thegreatforge/gokit/logger"
)

func main() {
	err := logger.Initialize()

	if err != nil {
		fmt.Println("failed to initialise logg")
		os.Exit(1)
	}

	err = config.Initialise(
		config.WithFiles([]string{"configs/config.yaml"}),
	)
	if err != nil {
		fmt.Printf("failed to init the config with error - %s", err.Error())
		os.Exit(1)
	}

	logger.Info("starting service: mychain")
}
