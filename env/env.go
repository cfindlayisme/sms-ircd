package env

import "os"

func GetServerName() string {
	return os.Getenv("SERVER_NAME")
}

func GetServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "6555"
	}

	return port
}
