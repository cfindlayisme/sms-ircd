package env

import "os"

func GetServerName() string {
	return os.Getenv("SERVER_NAME")
}
