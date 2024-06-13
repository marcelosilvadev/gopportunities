package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcelosilvadev/gopportunities.git/schemas"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error to create opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(context, "create-opening", opening)
}
