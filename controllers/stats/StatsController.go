package stats

import (
	"encoding/json"
	"mcs_service/db"
	u "mcs_service/utils"
	"net/http"
)

type EntityCountsResponse struct {
	TasksCount       uint  `json:"tasks_count"`
	SectionsCount    uint  `json:"sections_count"`
	MaterialsCount   uint  `json:"materials_count"`
	LessonsCount     uint  `json:"lessons_count"`
}

var GetEntityCount = func(w http.ResponseWriter, r *http.Request) {
	var tasksCount uint
	var sectionsCount uint
	var materialsCount uint
	var lessonsCount uint

	db := db.GetDB()
	db.Table("tasks").Where("deleted_at is null").Count(&tasksCount)
	db.Table("materials").Where("deleted_at is null").Count(&materialsCount)
	db.Table("sections").Where("deleted_at is null").Count(&sectionsCount)
	db.Table("lessons").Where("deleted_at is null").Count(&lessonsCount)

	response := EntityCountsResponse{
		TasksCount: tasksCount,
		SectionsCount: sectionsCount,
		MaterialsCount: materialsCount,
		LessonsCount: lessonsCount,
	}

	res, err := json.Marshal(response)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
