package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ojCezarFerreira/goportunities/schemas"
)

// @BasePath /api/v1
// @Summary Update Opening
// @Description Update a opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}
	c.BindJSON(&request)

	request.Validate()

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save Opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error updating opening")
	}

	sendSuccess(c, "update-opening", opening)
}
