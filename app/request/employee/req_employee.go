package employee

type ReqEmployee struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
