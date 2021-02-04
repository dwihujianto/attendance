package presence

import (
	"net/http"

	"attendance/app/model"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	o := []model.Presence{}
	model.DB.Find(&o)

	c.JSON(http.StatusOK, gin.H{"data": o})
}
