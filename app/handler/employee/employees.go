package employee

import (
	"net/http"

	"attendance/app/model"
	"github.com/gin-gonic/gin"
	"github.com/paulmach/go.geojson"

	req "attendance/app/request/employee"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func newGeoJson(point Point) ([]byte, error) {
	featureCollection := geojson.NewFeatureCollection()
	feature := geojson.NewPointFeature([]float64{point.Latitude, point.Longitude})
	featureCollection.AddFeature(feature)

	return featureCollection.MarshalJSON()
}

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
	r := req.ReqEmployee{}

	c.ShouldBindJSON(&r)

	point, _ := newGeoJson(Point{r.Latitude, r.Longitude})

	employee := model.Employee{
		Name:          r.Name,
		Email:         r.Email,
		LocationPoint: string(point),
	}

	model.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}
