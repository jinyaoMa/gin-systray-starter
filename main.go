package main

import (
	"flag"
	"io"
	"os"
	"this/server"
	"this/tray"

	"github.com/gin-gonic/gin"
)

var isDev = flag.Int("dev", 0, "set to development mode")
var hasTray = flag.Int("tray", 1, "set to enable system tray (bug when hot reload)")
var autoStart = flag.Int("server", 1, "set to auto start server")

func main() {
	flag.Parse()

	if *isDev == 1 {
		if *hasTray == 1 {
			tray.Run(func() {
				if *autoStart == 1 {
					tray.StartServerWithSwag()
				}
			})
		} else {
			server.Start(true)
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()

		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

		if *hasTray == 1 {
			tray.Run(func() {
				if *autoStart == 1 {
					tray.StartServer()
				}
			})
		} else {
			server.Start(false)
		}
	}
}
