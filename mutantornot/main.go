package main

import (
	"meli/mutantornot/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/mutant", handler.HandlePOSTMutant)
	router.Run(":8081")
}
