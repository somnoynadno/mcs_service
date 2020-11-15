package student

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetTasksBySectionID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sectionID := params["section_id"]

	password := r.FormValue("password")
	if password == "" {
		u.HandleBadRequest(w, errors.New("no password provided"))
		return
	}

	db := db.GetDB()

	section := entities.Section{}
	err := db.First(&section, sectionID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	if section.Password != nil && *section.Password != password {
		u.HandleForbidden(w, errors.New("wrong password"))
		return
	}

	var entities []entities.TaskForStudent
	err = db.Table("tasks").Preload("TaskType").Order("created_at ASC").
		Where("section_id = ?", sectionID).Find(&entities).Error

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	res, err := json.Marshal(entities)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
