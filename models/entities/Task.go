package entities

import (
	"mcs_service/models/auxiliary"
)

type Task struct {
	auxiliary.BaseModel
	Name        *string  `json:"name"              gorm:"not null"`
	Description string   `json:"description"`
	Solution    *string  `json:"solution"`
	Author      *string  `json:"author"`
	Difficulty  *int     `json:"difficulty"`
	SectionID	uint     `json:"section_id"`
	Section     *Section `json:"section,omitempty"`
}
