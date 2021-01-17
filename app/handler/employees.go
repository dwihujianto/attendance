package handler

import (
	"net/http"

	"github.com/dwihujianto/attendance/app/model"
	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	employees := []model.Employee{}
	model.DB.Find(&employees)

	c.JSON(http.StatusOK, gin.H{"data": employees})
}
