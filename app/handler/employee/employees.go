package employee

import (
	"net/http"

	"github.com/dwihujianto/attendance/app/model"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	employees := []model.Employee{}
	model.DB.Find(&employees)

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func GetById(c *gin.Context) {
	employee := model.Employee{}
	model.DB.First(&employee, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": employee})
}
