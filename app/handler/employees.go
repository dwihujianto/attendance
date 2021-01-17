package handler

import (
	//"encoding/json"
	"net/http"

	"github.com/dwihujianto/attendance/app/model"
	//"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)

	respondJSON(w, http.StatusOK, employees)
}