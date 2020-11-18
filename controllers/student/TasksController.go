package student

import (
	"crypto/aes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var GetTasksBySectionID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sectionID := params["section_id"]
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
		iv := "0000000000000000"
		key := u.PadKey(*section.Password, 32, "0")
		encrypted := u.AES256(string(res), key, iv, aes.BlockSize)

		log.Debug("section encrypted with key " + key + " and IV " + iv)
		u.Respond(w, u.Message(true, encrypted))
	}
}
