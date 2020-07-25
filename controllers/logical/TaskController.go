package logical

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetTasksBySectionID = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Task

	params := mux.Vars(r)
	sectionID := params["section_id"]

	db := db.GetDB()
	err := db.Preload("TaskType").Where("section_id = ?", sectionID).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var GetTasksByTaskTypeID = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Task

	params := mux.Vars(r)
	taskTypeID := params["task_type_id"]

	db := db.GetDB()
	err := db.Preload("TaskType").Where("task_type_id = ?", taskTypeID).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

