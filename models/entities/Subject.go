package entities

import "mcs_service/models/auxiliary"

type Subject struct {
	auxiliary.BaseModel
	Name        string     `json:"name"         gorm:"not null"`
	Description string     `json:"description"`
	Sections    []*Section `json:"sections,omitempty"`
	Teachers    *string    `json:"teachers"`
	// TODO: Rules       string     `json:"rules"`
}
