package main

import "log"

var App Application

func main() {
	err := App.load()
	if err != nil {
		log.Println(err)
		return
	}
	defer App.unload()

	App.start()
}
