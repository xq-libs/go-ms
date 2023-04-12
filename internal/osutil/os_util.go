package osutil

import "os"

func GetEnvValue(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}
