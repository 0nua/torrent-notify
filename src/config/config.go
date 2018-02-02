package config

import "os"

func GetToken() string {
	return os.Getenv("TG_BOT_TOKEN")
}

func GetDB() string {
	return os.Getenv("DB_PATH")
}