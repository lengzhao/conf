package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseDotENV parses the given data string and returns a map[string]string.
//
// The data string is split into lines, and each line is trimmed of leading and trailing spaces.
// If a line starts with "#" it is skipped.
// Otherwise, the line is split at the first occurrence of "=" or ":".
// If no "=" or ":" is found, the line is skipped and a log message is printed.
// The key and value are trimmed of leading and trailing spaces.
// The key-value pair is added to the map.
//
// Returns a map[string]string representing the parsed key-value pairs.
func ParseDotENV(data string) map[string]string {
	out := make(map[string]string)
	lines := strings.Split(data, "\n")
	for _, it := range lines {
		it = strings.TrimSpace(it)
		if strings.HasPrefix(it, "#") {
			continue
		}
		index := strings.IndexAny(it, "=:")
		if index == -1 || index >= len(it) {
			log.Println("unknow item:", it, index)
			continue
		}
		key := strings.TrimSpace(it[:index])
		value := strings.TrimSpace(it[index+1:])

		out[key] = value
	}

	return out
}

// ParseJson parses a JSON string and returns a map of string key-value pairs.
//
// It takes a single parameter:
//   - data: a string containing the JSON data to be parsed.
//
// It returns a map[string]string, which represents the parsed JSON data.
func ParseJson(data string) map[string]string {
	out := make(map[string]string)
	info := make(map[string]any)
	err := json.Unmarshal([]byte(data), &info)
	if err != nil {
		return nil
	}
	var ok bool
	for k, v := range info {
		out[k], ok = v.(string)
		if !ok {
			d, _ := json.Marshal(v)
			out[k] = string(d)
		}
	}
	return out
}

// ParseYML parses the given YAML data and returns a map[string]string.
//
// It takes a string `data` as a parameter and returns a map[string]string.
func ParseYML(data string) map[string]string {
	out := make(map[string]string)
	info := make(map[string]any)
	err := yaml.Unmarshal([]byte(data), &info)
	if err != nil {
		return nil
	}
	var ok bool
	for k, v := range info {
		out[k], ok = v.(string)
		if !ok {
			out[k] = fmt.Sprintf("%v", v)
		}
	}
	return out
}
