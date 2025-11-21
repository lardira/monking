package env

import (
	"log"
	"os"
)

func MustGetEnv(key string) string {
	gotEnv, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("didn't find env by key %v", key)
	}

	return gotEnv
}
