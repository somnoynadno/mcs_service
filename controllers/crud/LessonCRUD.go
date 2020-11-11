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

var LessonCreate = func(w http.ResponseWriter, r *http.Request) {
	Lesson := &entities.Lesson{}
	err := json.NewDecoder(r.Body).Decode(Lesson)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(Lesson).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(Lesson)
		u.RespondJSON(w, res)
	}
}

var LessonRetrieve = func(w http.ResponseWriter, r *http.Request) {
	Lesson := &entities.Lesson{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("LessonType").First(&Lesson, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(Lesson)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var LessonUpdate = func(w http.ResponseWriter, r *http.Request) {
	Lesson := &entities.Lesson{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Lesson, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newLesson := &entities.Lesson{}
	err = json.NewDecoder(r.Body).Decode(newLesson)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&Lesson).Updates(newLesson).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LessonDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.Lesson{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}
