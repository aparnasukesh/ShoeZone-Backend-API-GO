package main

import (
	"os"

	"github.com/aparnasukesh/shoezone/pkg/config"
	"github.com/aparnasukesh/shoezone/pkg/server"
)

func main() {

	config.LoadConfig()
	engine := server.SeverConnect()
	engine.Run(":" + os.Getenv("PORT"))
}
