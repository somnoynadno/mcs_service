package entities

import (
	"mcs_service/models/auxiliary"
)

type Task struct {
	auxiliary.BaseModel
	Name           string    `json:"name"              gorm:"not null"`
	Description    string    `json:"description"`
	Solution       *string   `json:"solution"`
	Author         *string   `json:"author"`
	SolutionAuthor *string   `json:"solution_author"`
	Difficulty     *int      `json:"difficulty"`
	SectionID	   uint      `json:"section_id"`
	Section        *Section  `json:"section,omitempty"`
	TaskTypeID     uint      `json:"task_type_id"`
	TaskType       *TaskType `json:"task_type,omitempty"`
	Notes          *string   `json:"notes"`
	Lessons        []*Lesson `json:"lessons,omitempty" gorm:"many2many:task_lessons;"`
}

type TaskForStudent struct {
	auxiliary.BaseModel
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Difficulty     *int      `json:"difficulty"`
	SectionID	   uint      `json:"section_id"`
	Section        *Section  `json:"section,omitempty"`
	TaskTypeID     uint      `json:"task_type_id"`
	TaskType       *TaskType `json:"task_type,omitempty"`
}
