package conf

import (
	"os"
	"strconv"
)

func Get(key, defaultValue string) string {
	out := localConf[key]
	if len(out) > 0 {
		return out
	}
	out = os.Getenv(key)
	if len(out) > 0 {
		return out
	}
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetFloat(key string, defaultValue float64) float64 {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultValue
	}
	return v
}
