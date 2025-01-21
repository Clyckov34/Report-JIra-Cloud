package main

import (
	"log"
	"os"
	"report/internal"
	"report/pkg/config"
)

var params = &config.Config{}

func init() {
	params = config.GetConfig()
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("укажите параметры (Флаги). Подробно: --help")
	} else {
		if err := internal.GetReport(params); err != nil {
			log.Fatalln(err)
		}
	}
}
