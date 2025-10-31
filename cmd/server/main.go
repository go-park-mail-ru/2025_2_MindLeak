package main

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/apiserver"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/server"
	_ "github.com/go-park-mail-ru/2025_2_MindLeak/swagger"
)

func main() {
	config := server.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", config)
	if err != nil {
		logger.Error(nil, err.Error(), nil)
		log.Fatal(err)
	}
	s, err := apiserver.New(config)
	if err != nil {
		logger.Error(nil, err.Error(), nil)
		log.Fatal(err)
	}
	s.StartServer()
}
