package employee

import (
	"net/http"

	"attendance/app/model"
	"github.com/gin-gonic/gin"

	req "attendance/app/request/employee"
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

func CreateNew(c *gin.Context) {
	var r req.ReqEmployee
	c.ShouldBindJSON(&r)

	employee := model.Employee{
		Name:          r.Name,
		Email:         r.Email,
		LocationPoint: r.LocationPoint,
	}

	model.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}
