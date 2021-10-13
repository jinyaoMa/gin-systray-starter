package main

import (
	"flag"
	"this/server"
	"this/tray"

	"github.com/gin-gonic/gin"
)

var isDev = flag.Int("dev", 0, "set to development mode")

func main() {
	flag.Parse()

	if *isDev == 1 {
		server.Run()
	} else {
		gin.SetMode(gin.ReleaseMode)
		tray.Run()
	}
}
