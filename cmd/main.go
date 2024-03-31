package main

import (
	"fmt"
	"os"

	"github.com/aparnasukesh/shoezone/pkg/config"
	"github.com/aparnasukesh/shoezone/pkg/controller"
	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/server"
)

func main() {

	config.LoadConfig()
	db.DbConnect()
	engine := server.SeverConnect()
	controller.Routes(engine)
	fmt.Println("Running...")
	engine.Run(":" + os.Getenv("PORT"))
}
