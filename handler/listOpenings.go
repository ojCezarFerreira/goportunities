package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ojCezarFerreira/goportunities/schemas"
)

// @BasePath /api/v1
// @Summary List Openings
// @Description List all Openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
