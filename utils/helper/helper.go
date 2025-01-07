package helper

import "os"

func GetENV(key, defaultValue string) string {
	getEnv := os.Getenv(key)
	// fmt.Println("ENV: ", getEnv)
	if len(getEnv) == 0 || getEnv == "" {
		return defaultValue
	}
	return getEnv
}
