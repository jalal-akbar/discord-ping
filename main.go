package main

import (
	"fmt"

	"github.com/jalal-akbar/discord-ping/bot"
	"github.com/jalal-akbar/discord-ping/config"
)

func main() {
	// read config
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
}
