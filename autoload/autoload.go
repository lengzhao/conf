package autoload

import (
	"log"

	"github.com/lengzhao/conf"
)

func init() {
	fn := conf.LoadOneToENV()

	if fn == "" {
		log.Println("not found any config file")
	}
}
