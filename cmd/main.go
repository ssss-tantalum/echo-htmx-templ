package main

import (
	"fmt"

	"github.com/ssss-tantalum/echo-templ-htmx/internal/core"
	"github.com/ssss-tantalum/echo-templ-htmx/internal/server"
)

func main() {
	config, err := core.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("🚀 Http server started on %s:%d ⭐\n", config.ServerHost, config.SeverPort)
	errServ := server.RunServer()
	if errServ != nil {
		panic(errServ)
	}
}
