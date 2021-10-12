package server

import (
	"context"
	"log"
	"net/http"
	"this/server/routers"
	"time"

	"github.com/gin-gonic/gin"
)

var server *http.Server

func init() {
	r := gin.Default()

	routers.TestRouter(r.Group("/"))

	server = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}

func Run() {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
