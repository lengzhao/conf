package conf

import "os"

func SaveToENV(e map[string]string, cover bool) {
	for k, v := range e {
		if !cover {
			if os.Getenv(k) != "" {
				continue
			}
		}
		os.Setenv(k, v)
	}
}

var localConf map[string]string

func SaveToMemory(e map[string]string, cover bool) {
	for k, v := range e {
		if !cover {
			if localConf[k] != "" {
				continue
			}
		}
		localConf[k] = v
	}
}
