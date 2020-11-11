package entities

import "mcs_service/models/auxiliary"

type LessonType struct {
	auxiliary.BaseModelIgnore
	Name string `json:"name" gorm:"not null"`
}

var DefaultLessonTypes = []string {
	"main",
	"homework",
	"additional",
	"exam",
}