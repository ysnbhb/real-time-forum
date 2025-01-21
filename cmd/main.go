package main

import (
	"log"

	"real-time-forum/module"
)

func main() {
	_, err := module.Init()
	if err != nil {
		log.Fatal(err)
	}
}
