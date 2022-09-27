package main

import (
	"github.com/Phazon85/restapp-demo/services/httprouter"
)

func main() {
	//Create Gin Router with Handlers attached
	r := httprouter.New()

	//Run Gin server
	r.Run()
}
