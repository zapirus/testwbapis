package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/zapirus/testwbapis/internal/handlers"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "")
}

func main() {
	flag.Parse()
	config := handlers.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}
	s := handlers.New(config)

	s.Run()

}
