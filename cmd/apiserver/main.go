package main

import (
	"flag"
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
	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Parsing flag")
	flag.Parse()
	l.Infoln("Config initialization")
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Starting apiserver addr : %s\n", config.BindAddr)
	if err := apiserver.Start(config); err != nil {
		l.Fatal(err)
	}
}
