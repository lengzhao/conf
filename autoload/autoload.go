package autoload

import "github.com/lengzhao/conf"

func init() {
	conf.Load(func(e map[string]string, cover bool) {
		conf.SaveToMemory(e, true)
		conf.SaveToENV(e, true)
	}, true, ".env", "conf.json")
}
