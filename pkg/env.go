package env

import (
	"os"
)

// GetEnvOrDefault returns the environment variable value related to the
// given key if this one exists, otherwise it returns the default passed value.
func GetEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if "" == v {
		return def
	}
	return v
}
