package main

import (
	Elysium "github.com/hipinion/elysium/src"
	"log"
)

const (
	VERSION = "0.01"
	CONFIG  = "config/config.json"
)

func init() {
	log.Println("\360\237\215\224 \tElysium v" + VERSION)
}

func main() {
	Elysium.Init(CONFIG)
	Elysium.Serve()
}
