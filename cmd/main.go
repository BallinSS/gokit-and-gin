package main

import (
	"log"
	"net/http"

	"gin-gokit-microservice/pkg/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 设置路由和处理器
	transport.SetupRoutes(router)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
