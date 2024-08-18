package conf

import (
	"os"
	"strconv"
	"strings"
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

func GetInt64(key string, defaultValue int64) int64 {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetUint(key string, defaultValue uint) uint {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint(v)
}

func GetUint64(key string, defaultValue uint64) uint64 {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetFloat(key string, defaultValue float32) float32 {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return defaultValue
	}
	return float32(v)
}

func GetFloat64(key string, defaultValue float64) float64 {
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

func GetBool(key string, defaultValue bool) bool {
	str := Get(key, "")
	if len(str) == 0 {
		return defaultValue
	}
	str = strings.ToLower(str)
	switch str {
	case "true", "1":
		return true
	case "false", "0":
		return false
	}
	return defaultValue
}
