package entities

import "mcs_service/models/auxiliary"

type Section struct {
	auxiliary.BaseModel
	Name          string       `json:"name"                   gorm:"not null"`
	Description   string       `json:"description"`
	Password      *string      `json:"password,omitempty"`
	SubjectID     uint         `json:"subject_id"`
	Subject       *Subject     `json:"subject,omitempty"`
	Tasks         []*Task      `json:"tasks,omitempty"`
	SectionTypeID uint         `json:"section_type_id"`
	SectionType   *SectionType `json:"section_type,omitempty"`
	Materials     []*Material  `json:"materials,omitempty"`
}
