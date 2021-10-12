package main

import (
	"os"
	"this/server"
)

func main() {
	if os.Getenv("IS_PRODUCTION") == "1" {

	} else {
		s1 := server.New()
		s1.Run()
	}
}
