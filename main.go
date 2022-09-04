package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpAddr = ":3000"

func main() {
	fmt.Println("Server running on", httpAddr)

	srv := gin.New()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpAddr))
}

func healthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Everything is ok!")
}
