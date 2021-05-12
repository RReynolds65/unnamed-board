package handler

import (
	"net/http"

	"api/app/model"

	"github.com/jinzhu/gorm"
)

func GetAllPosts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	posts := []model.Employee{}
	db.Find(&posts)
	respondJSON(w, http.StatusOK, posts)
}
