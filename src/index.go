package main

import (
	"tgbot"
	"config"
)

func main() {
	if (config.IsLoaded()) {
		tgbot.Init()
	}
}
