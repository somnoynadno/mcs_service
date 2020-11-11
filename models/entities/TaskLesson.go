package entities

import "mcs_service/models/auxiliary"

type TaskLesson struct {
	auxiliary.BaseModelIgnore
	TaskID   uint  `json:"task_id"`
	LessonID uint  `json:"lesson_id"`
}
