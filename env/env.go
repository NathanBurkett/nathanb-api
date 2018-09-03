package env

import (
	"os"
	"log"
)

// Must Get the env key's value. Log error and exit if doesn't exist
func Must(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env variable %s does not exist", key)
	}

	return value
}
