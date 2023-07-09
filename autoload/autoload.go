package autoload

import (
	"log"

	"github.com/lengzhao/conf"
)

func init() {
	fn := conf.LoadAnyOne(func(e map[string]string, cover bool) {
		conf.SaveToMemory(e, true)
		conf.SaveToENV(e, true)
	}, true, ".env", "conf.json", "config.json", "conf.yml", "config.yml")

	if fn == "" {
		log.Println("not found any config file")
	}
}
