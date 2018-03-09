package config

import (
	"os"
	"strconv"
	"error"
	"path/filepath"
	"github.com/joho/godotenv"
)

var configLoaded = initConfig()

func IsLoaded() bool {
	return configLoaded
}

func GetToken() string {
	return os.Getenv("TG_BOT_TOKEN")
}

func GetDB() string {
	path, err := filepath.Abs(os.Getenv("DB_PATH"))
	error.Catch(err)
	return path
}

func GetSaveTimeout() float64 {
	timeout, err := strconv.ParseFloat(os.Getenv("SAVE_TIMEOUT"), 64)
	error.Catch(err)
	return timeout
}

func GetRutrackerApi() string {
	return os.Getenv("RUTRACKER_API")
}

func initConfig() bool {
	err := godotenv.Load(".env")
	return err == nil;
}