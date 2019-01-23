package main

import (
	"meli/dnastats/business"
	"meli/dnastats/handler"
	"meli/rabbit"

	"github.com/gin-gonic/gin"
)

func main() {
	// Trigger the rabbitmq consumer
	go rabbit.Consume(business.SaveDNAIfNotExists)

	router := gin.Default()

	router.GET("/stats", handler.HandleGETStats)
	router.Run(":8082")
}
