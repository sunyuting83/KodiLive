package main

import (
	"KodiLive/router"
	"flag"
	"strings"
)

func main() {
	var (
		port  string
		proxy int
		onoff bool = false
	)
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.IntVar(&proxy, "proxy", 0, "端口号，默认为3000")
	flag.Parse()
	if proxy > 0 {
		onoff = true
	}
	app := router.InitRouter(onoff)

	app.Listen(strings.Join([]string{":", port}, ""))
}
