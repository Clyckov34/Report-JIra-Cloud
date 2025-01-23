package main

import (
	"log"
	"report/internal"
	"report/pkg/config"
)

var params = &config.Config{}

func init() {
	params = config.GetConfig()
}

func main() {
	if err := internal.GetReport(params); err != nil {
		log.Fatalln(err)
	}
}
