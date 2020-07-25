package entities

import "mcs_service/models/auxiliary"

type Material struct {
	auxiliary.BaseModel
	Name        string   `json:"name"              gorm:"not null"`
	Content     string   `json:"content"`
	IsVisible   bool     `json:"is_visible"        gorm:"not null;default:true"`
	SectionID   uint     `json:"section_id"`
	Section     *Section `json:"section,omitempty"`
}
