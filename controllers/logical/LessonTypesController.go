package logical

import (
	"encoding/json"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetAllLessonTypes = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.LessonType

	db := db.GetDB()
	err := db.Order("created_at ASC").Find(&entities).Error

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

