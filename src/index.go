package main

import (
	"error"
	"tgbot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env");
	error.Check(err)
	tgbot.Init()
}
