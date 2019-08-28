package main

import (
	"flag"
	"fmt"
	"log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "service config")
}

func main() {

	flag.Parse()

	if len(configFile) == 0 {
		configFile = "config/config.yml"
	}

	// load app config
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("Load configuration %s error: %s", configFile, err)
	}

	// load log config
	logger, err:= getLogger(&config.Logger)
	if err != nil {
		log.Fatalf("Logger setup error: %s", err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	logger.Debug(fmt.Sprintf("Config %#v\n", config))
}
