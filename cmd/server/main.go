package main

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/server"
	_ "github.com/go-park-mail-ru/2025_2_MindLeak/swagger"
)

func main() {
	server.StartServer()
}
