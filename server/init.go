package server

import (
	"context"
	"log"
	"net/http"
	"this/server/routers"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var server *http.Server

var IsRunning = false

func Start() {
	router = gin.Default()

	routers.TestRouter(router.Group("/"))

	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func Run() {
	if IsRunning {
		return
	}
	IsRunning = true
	log.Println("Server Start!")

	go Start()
}

func Stop() {
	if !IsRunning {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server Stop!")
	IsRunning = false
}
