package config

import (
	"fmt"
	"os"
)

func Get() string {
	bind := os.Getenv("BIND")
	port := os.Getenv("PORT")

	if bind == "" {
		bind = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf("%s:%s", bind, port)
}
