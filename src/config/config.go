package config

import (
	"os"
	"strconv"
	"fail"
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
	fail.Catch(err)
	return path
}

func GetSaveTimeout() float64 {
	timeout, err := strconv.ParseFloat(os.Getenv("SAVE_TIMEOUT"), 64)
	fail.Catch(err)
	return timeout
}

func GetRutrackerApi() string {
	return os.Getenv("RUTRACKER_API")
}

func GetUpdateTimeout() int {
	timeout, err := strconv.Atoi(os.Getenv("UPDATE_TIMEOUT"));
	fail.Catch(err)
	return timeout
}

func initConfig() bool {
	err := godotenv.Load(".env")
	return err == nil;
}