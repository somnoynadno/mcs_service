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

var GetEncryptedTasksBySectionID = func(w http.ResponseWriter, r *http.Request) {
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
		if section.Password == nil {
			log.Info("no password for section provided")
			p := ""
			section.Password = &p
		}

		iv := "0000000000000000"
		key := u.PadKey(*section.Password, 32, "0")
		encrypted := u.AES256(string(res), key, iv, aes.BlockSize)

		log.Debug("section encrypted with key " + key + " and IV " + iv)
		u.Respond(w, u.Message(true, encrypted))
	}
}

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

	if section.Password == nil {
		log.Info("no password for section provided")
		p := ""
		section.Password = &p
	}

	iv := "0000000000000000"
	key := u.PadKey(*section.Password, 32, "0")

	for i, v := range entities {
		d := v.Description
		encrypted := u.AES256(d, key, iv, aes.BlockSize)
		entities[i].Description = encrypted
	}
	log.Debug("descriptions encrypted with key " + key + " and IV " + iv)

	res, err := json.Marshal(entities)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var GetAllTasks = func(w http.ResponseWriter, r *http.Request) {
	var entities []entities.TaskForStudent

	db := db.GetDB()
	err := db.Table("tasks").Preload("TaskType").Preload("Section").Find(&entities).Error

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	for i, v := range entities {
		password := v.Section.Password
		if password == nil {
			log.Debug("no password for section provided")
			p := ""
			password = &p
		}

		iv := "0000000000000000"
		key := u.PadKey(*password, 32, "0")

		d := v.Description
		encrypted := u.AES256(d, key, iv, aes.BlockSize)

		entities[i].Description = encrypted
		entities[i].Section = nil // hide section
	}

	log.Debug("descriptions encrypted successfully")
	res, err := json.Marshal(entities)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}