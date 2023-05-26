package conf

import (
	"log"
	"strings"

	"github.com/buger/jsonparser"
)

func ParseDotENV(data string) map[string]string {
	out := make(map[string]string)
	lines := strings.Split(data, "\n")
	for _, it := range lines {
		it = strings.TrimSpace(it)
		if strings.HasPrefix(it, "#") {
			continue
		}
		if strings.Contains(it, "=") {
			sp := strings.Split(it, "=")
			key := strings.TrimSpace(sp[0])
			value := strings.TrimSpace(strings.Join(sp[1:], "="))
			out[key] = value
			continue
		}
		if strings.Contains(it, ":") {
			sp := strings.Split(it, ":")
			key := strings.TrimSpace(sp[0])
			value := strings.TrimSpace(strings.Join(sp[1:], ":"))
			out[key] = value
			continue
		}
		log.Println("unknow item:", it)
	}

	return out
}

func ParseJson(data string) map[string]string {
	out := make(map[string]string)
	jsonparser.ObjectEach([]byte(data), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		// fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
		out[string(key)] = string(value)
		return nil
	})

	return out
}
