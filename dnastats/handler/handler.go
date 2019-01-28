package handler

import (
	"log"
	"meli/dnastats/business"
	"meli/dnastats/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGETStats handles a request to a stats service
func HandleGETStats(ctx *gin.Context) {
	mutants, humans, ratio := business.GetDNAStats()
	statCon := container.Stat{
		Mutants: mutants,
		Humans:  humans,
		Ratio:   ratio,
	}
	log.Printf("Stats: %d, %d, %f", mutants, humans, ratio)
	ctx.JSON(http.StatusOK, statCon)

}
