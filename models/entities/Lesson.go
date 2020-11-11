package entities

import (
	"mcs_service/models/auxiliary"
	"time"
)

type Lesson struct {
	auxiliary.BaseModel
	Name           string      `json:"name"                   gorm:"not null"`
	AdditionalInfo *string     `json:"additional_info"`
	Password       *string     `json:"password"`
	FromDate       *time.Time  `json:"from_date"`
	DueDate        *time.Time  `json:"due_date"`
	IsVisible      bool        `json:"is_visible"             gorm:"not null;default:true"`
	LessonTypeID   uint        `json:"lesson_type_id"`
	LessonType     *LessonType `json:"lesson_type,omitempty"`
	SubjectID      *uint       `json:"subject_id"`
	Subject        *Subject    `json:"subject,omitempty"`
	Tasks          []*Task     `json:"tasks,omitempty"        gorm:"many2many:task_lessons;"`
}
