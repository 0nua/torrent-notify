package config

var config = map[string]string {
	"token": "483400559:AAFf-QjzbN-svlQmVpUIrFvW6oe4TA9wI0k",
	"db": "db/users.json",
}

func GetToken() string {
	return config["token"]
}

func GetDB() string {
	return config["db"]
}