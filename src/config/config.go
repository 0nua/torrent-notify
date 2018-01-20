package config

var config = map[string]string {
	"token": "123",
}

func GetToken() string {
	return config["token"]
}