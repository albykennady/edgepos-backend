package config

import "os"

type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("PORT")

	// TODO: If port is empty, assign the default value "8080".
	if port == "" {
		port = "8080"
	}

	return Config{
		Port: port,
	}
}
