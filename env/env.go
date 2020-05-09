package env

import "os"

// GetEnvOrDefault ges the env var value related to the given key,
// if it does not exists, it return the given default value.
func GetEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if "" == v {
		return def
	}
	return v
}
