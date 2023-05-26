package conf

import (
	"log"
	"os"
	"path"
)

type LoadCB func(e map[string]string, cover bool)

func Load(cb LoadCB, cover bool, files ...string) {
	for _, fn := range files {
		data, err := os.ReadFile(fn)
		if err != nil {
			log.Println("not found file:", fn, err)
			continue
		}
		var items map[string]string
		switch path.Ext(fn) {
		case ".env":
			items = ParseDotENV(string(data))
		case ".json":
			items = ParseJson(string(data))
		}
		cb(items, true)
	}
}
