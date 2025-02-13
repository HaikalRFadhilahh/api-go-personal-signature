package pkg

import "os"

func GetEnvValue(k, d string) string {
	if os.Getenv(k) == "" {
		return d
	}

	return os.Getenv(k)
}
