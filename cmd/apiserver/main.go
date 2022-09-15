package main

import (
	"flag"
	"log"
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

func main() {
	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Parsing flag")
	flag.Parse()
	l.Infoln("Config initialization")
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	l.Infof("Starting apiserver addr : %s\n", config.BindAddr)
	if err := apiserver.Start(config); err != nil {
		l.Fatal(err)
	}
}
