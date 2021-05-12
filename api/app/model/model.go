package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	post_id    int
	image_file string
	user       string
	date       string
	thread_id  int
	post_text  string
	Status     bool
}

func (p *Post) Enable() {
	p.Status = true
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Post{})
	return db
}
