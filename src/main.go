package main

import (
	"log"

	server "github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/server"
)

func main() {

	log.Print("INFO : starting the application ....")

	server.LoadRoute()

}
