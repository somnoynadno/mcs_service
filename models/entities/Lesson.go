package entities

import "time"

type Lesson struct {
	Name         string               `json:"name"                  gorm:"not null"`
	FromDate     time.Time            `json:"from_date"`
	DueDate      time.Time            `json:"due_date"`
	Password     string               `json:"password"              gorm:"not null"`
	IsVisible    bool                 `json:"is_visible"            gorm:"not null;default:true"`
	LessonTypeID *uint                `json:"lesson_type_id"`
	LessonType   *LessonType          `json:"lesson_type,omitempty"`
	TaskIDs      []*uint              `json:"task_ids,omitempty"    gorm:"-"`
	Tasks        []*TaskForStudents   `json:"tasks,omitempty"`
}
