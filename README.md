# conf

load config file to os env.

support files(anyone):

1. .env
2. conf.json
3. config.json
4. conf.yml
5. config.yml

you can get configuration by os.Getenv

~~~golang
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
~~~

## warning

If your app is a macos app, the file cannot be found during initialization, so you need to actively load it at the main function.

~~~golang
package main

import (
    "log"
    "os"

    "github.com/lengzhao/conf"
)

func main() {
    //conf.Load(".env","conf.json")
    //conf.Load("a.env")
    conf.Load()// anyone of [".env", "conf.json", "config.json", "conf.yml", "config.yml"]
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
~~~
