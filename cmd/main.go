package main

import (
	"log"

	"github.com/ramses2099/cleanarchitecturewithgo/api"
)

const port string = "3001"

func main() {
	log.Println("starting API cmd")
	api.Start(port)
}
