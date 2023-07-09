package main

import (
	"log"
	"os"

	"github.com/lengzhao/conf"
	_ "github.com/lengzhao/conf/autoload"
)

func main() {
	{
		val := os.Getenv("TestKey1")
		if val != "TestValue1" {
			log.Panic("fail to load .env")
		}
		log.Println("TestKey1:", val)

		val2 := os.Getenv("TestInt")
		if val2 != "100" {
			log.Panic("fail to load config", val2)
		}
	}
	{
		conf.Load(conf.SaveToENV, true, "conf.json")
		v1 := os.Getenv("JsonKey")
		if v1 != "JsonValue" {
			log.Panic("fail to load conf.json")
		}
		v2 := os.Getenv("JsonInt")
		if v2 != "100" {
			log.Panic("fail to load conf.json")
		}
	}
}
