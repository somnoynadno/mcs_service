package logical

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetMaterialsBySectionID = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Material

	params := mux.Vars(r)
	sectionID := params["section_id"]

	db := db.GetDB()
	err := db.Where("section_id = ?", sectionID).Find(&entities).Error

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
