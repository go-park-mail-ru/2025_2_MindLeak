package main

import (
	"github.com/BurntSushi/toml"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/apiserver"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/server"
	_ "github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	_ "github.com/go-park-mail-ru/2025_2_MindLeak/swagger"
	"log"
)

func main() {
	config := server.NewConfig()
	_, err := toml.DecodeFile("configs/server.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	s, err := apiserver.New(config)
	if err != nil {
		log.Fatal(err)
	}
	s.StartServer()
}
