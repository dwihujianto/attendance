package employee

type ReqEmployee struct {
	Name          string `form:"name"`
	Email         string `form:"email"`
	LocationPoint string `form:location_point`
}
