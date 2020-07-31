package entities

import "mcs_service/models/auxiliary"

type TaskType struct {
	auxiliary.BaseModelIgnore
	Name string `json:"name" gorm:"not null"`
}

var DefaultTaskTypes = []string {
	"suggestion", // new task suggestion
	"task",       // existing task in module
	"homework",   // task as a homework
	"additional", // additional tasks
	"control",    // control tasks
	"rejected",   // declined suggestions
}
