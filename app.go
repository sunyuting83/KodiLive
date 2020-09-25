package main

import (
	router "KodiLive/router"
	"flag"
	"strings"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()

	app := router.InitRouter()

	app.Listen(strings.Join([]string{":", port}, ""))
}
