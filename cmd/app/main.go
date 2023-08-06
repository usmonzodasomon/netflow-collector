package main

import (
	"log"
	"udp/config"
	"udp/internal/app"
)

func main() {
	if err := config.GetConfigs(); err != nil {
		log.Fatal("error while getting configs: ", err.Error())
	}

	app.Run()
}
