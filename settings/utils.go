package settings

import "os"

func getEnvironment(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return def
}
