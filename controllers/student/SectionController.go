package student

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetSectionsBySubjectID = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.SectionForStudent

	params := mux.Vars(r)
	subjectID := params["subject_id"]

	db := db.GetDB()
	err := db.Table("sections").Preload("SectionType").Order("created_at ASC").
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

