package main

import (
	"log"

	"./routers"

	"github.com/urfave/negroni"
)

func main() {
	router := routers.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening:")
	n.Run(":3001")
}
