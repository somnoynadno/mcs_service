package entities

import "mcs_service/models/auxiliary"

type Section struct {
	auxiliary.BaseModel
	Name        string   `json:"name"              gorm:"not null"`
	Description string   `json:"description"`
	SubjectID   uint     `json:"subject_id"`
	Subject     *Subject `json:"subject,omitempty"`
	Tasks       []*Task  `json:"tasks,omitempty"`
}
