package handler

import (
	"meli/mutantornot/container"
	"meli/mutantornot/util"
	"meli/rabbit"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlePOSTMutant receive POST call to /mutant/
func HandlePOSTMutant(ctx *gin.Context) {
	var dna container.DnaMatrix
	if err := ctx.ShouldBindJSON(&dna); err != nil {
		ctx.JSON(http.StatusNotFound, "Invalid request")
		return
	}

	isMutant, err := util.IsMutant(dna.Dna)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "The DNA matrix is invalid")
		return
	}

	if isMutant {
		ctx.JSON(http.StatusOK, "The DNA is a mutant one")
	} else {
		ctx.JSON(http.StatusForbidden, "The DNA is NOT mutant")
	}

	// Publish the message in async mode
	go publishRabbitMessage(isMutant, dna)
}

func publishRabbitMessage(isMutant bool, dna container.DnaMatrix) {
	var message rabbit.DnaMessage
	message.Dna = dna.Dna
	message.IsMutant = isMutant
	rabbit.Publish(message)
}
