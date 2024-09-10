package util

import "os"

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func GetSlideDir() string {
	return getEnv("SLIDE_DIR", ".")
}

func GetPort() string {
	return getEnv("PORT", "8888")
}
