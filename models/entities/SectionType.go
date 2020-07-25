package entities

import "mcs_service/models/auxiliary"

type SectionType struct {
	auxiliary.BaseModelIgnore
	Name string `json:"name" gorm:"not null"`
}

var DefaultSectionTypes = []string {
	"suggestion",   // new section suggestions
	"main",         // existing section in subject
	"experimental", // section marked as experimental
	"rejected",     // declined suggestions
}
