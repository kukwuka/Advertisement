package main

import (
	server "Advertising/internal/app/api"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "config/config.toml", "path to config file (toml)")
}

func main() {
	flag.Parse()
	config := server.NewConfig()
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
