package conf

import "os"

func SaveToENV(e map[string]string, cover bool) {
	for k, v := range e {
		nv := os.Getenv(k)
		if nv == "" || cover {
			nv = v
		}
		os.Setenv(k, nv)
	}
}

var localConf map[string]string

func init() {
	localConf = make(map[string]string)
}

func SaveToMemory(e map[string]string, cover bool) {
	for k, v := range e {
		if _, ok := localConf[k]; !ok || cover {
			localConf[k] = v
		}
	}
}
