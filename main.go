package main

import (
	"flag"
	"log"

	"github.com/Uzumachi02/link2.xyz/server"
	"github.com/Uzumachi02/link2.xyz/store/postgresql"
	"github.com/kayac/go-config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/default.toml", "path to config file")
}

func main() {
	flag.Parse()

	conf := server.NewConfig()
	err := config.LoadTOML(conf, configPath)
	if err != nil {
		log.Fatal(err)
	}

	postgresql.NewStore(conf.DatabaseURL)

	if err := server.Start(conf); err != nil {
		log.Fatal(err)
	}
}
