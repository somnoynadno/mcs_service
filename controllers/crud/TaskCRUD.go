package crud

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var TaskCreate = func(w http.ResponseWriter, r *http.Request) {
	Task := &entities.Task{}
	err := json.NewDecoder(r.Body).Decode(Task)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(Task).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(Task)
		u.RespondJSON(w, res)
	}
}

var TaskRetrieve = func(w http.ResponseWriter, r *http.Request) {
	Task := &entities.Task{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("TaskType").First(&Task, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(Task)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var TaskUpdate = func(w http.ResponseWriter, r *http.Request) {
	Task := &entities.Task{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Task, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newTask := &entities.Task{}
	err = json.NewDecoder(r.Body).Decode(newTask)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&Task).Updates(newTask).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var TaskDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.Task{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}