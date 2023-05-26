package main

import (
	"log"
	"os"

	_ "github.com/lengzhao/conf/autoload"
)

func main() {
	val := os.Getenv("TestKey1")
	if val != "TestValue1" {
		log.Panic("fail to load .env")
	}
	log.Println("TestKey1:", val)

	val2 := os.Getenv("JsonKey")
	if val2 != "JsonValue" {
		log.Panic("fail to load conf.json")
	}
	log.Println("JsonKey:", val2)
	log.Println("JsonInt:", os.Getenv("JsonInt"))
}
