package env

import (
	"os"
)

func GetEnvOrDefault(key, defaulfVal string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaulfVal
	}

	return val
}
