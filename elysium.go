package main

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	VERSION = "0.01"
	CONFIG  = "config/config.json"
)

type Config struct {
}

func loadConfig() {
	cfile, err := os.Open(CONFIG)
	if err != nil {
		log.Fatalln("Could not open config", CONFIG)
	}
	config, err := ioutil.ReadAll(cfile)
	log.Println(config)
}

func init() {
	log.Println("\360\237\215\224 \tElysium v" + VERSION)
}

func main() {

}
