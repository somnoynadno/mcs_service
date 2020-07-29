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
}

var GetEntityCount = func(w http.ResponseWriter, r *http.Request) {
	var tasksCount uint
	var sectionsCount uint
	var materialsCount uint

	db := db.GetDB()
	db.Table("tasks").Count(&tasksCount)
	db.Table("materials").Count(&materialsCount)
	db.Table("sections").Count(&sectionsCount)

	response := EntityCountsResponse{
		TasksCount: tasksCount,
		SectionsCount: sectionsCount,
		MaterialsCount: materialsCount,
	}

	res, err := json.Marshal(response)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
