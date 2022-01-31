package main

import (
	"os"

	"github.com/marmotedu/goserver/internal/goserver"
)

// @title Apiserver Example API
// @version 1.0
// @description apiserver demo

// @contact.name lkong
// @contact.url http://www.swagger.io/support
// @contact.email 466701708@qq.com

// @host localhost:8080
// @BasePath /v1.
func main() {
	command := goserver.NewGoServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
