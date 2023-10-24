package main

import (
	"flag"
	"fmt"
	"restApi/internal/app/apiserver"
	"restApi/internal/app/logging"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// @title PVSystem24 API
// @version 1.12.0
// @description Swag documentaion for PVSystem24 API

// @host localhost:4000
// @BasePath /

func main() {

	logger, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	config := apiserver.NewConfig()
	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		logger.Error("error while decoding config file ", err)
		panic(err)
	}

	logger.Info(fmt.Sprintf("Starting apiserver addr : %s\n", config.BindAddr))
	if err := apiserver.Start(config, logger); err != nil {
		panic(err)
	}
}
