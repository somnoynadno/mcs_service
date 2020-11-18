package logical

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetSectionsBySubjectID = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Section

	params := mux.Vars(r)
	subjectID := params["subject_id"]

	db := db.GetDB()
	err := db.Preload("SectionType").Order("created_at ASC").
		Where("subject_id = ?", subjectID).Find(&entities).Error

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

var GetAllSections = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Section

	db := db.GetDB()
	err := db.Preload("SectionType").Order("created_at ASC").Find(&entities).Error

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