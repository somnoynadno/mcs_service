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

var MaterialCreate = func(w http.ResponseWriter, r *http.Request) {
	Material := &entities.Material{}
	err := json.NewDecoder(r.Body).Decode(Material)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(Material).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(Material)
		u.RespondJSON(w, res)
	}
}

var MaterialRetrieve = func(w http.ResponseWriter, r *http.Request) {
	Material := &entities.Material{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Material, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(Material)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var MaterialUpdate = func(w http.ResponseWriter, r *http.Request) {
	Material := &entities.Material{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Material, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newMaterial := &entities.Material{}
	err = json.NewDecoder(r.Body).Decode(newMaterial)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db.Model(&Material).Update("is_visible", newMaterial.IsVisible)
	err = db.Model(&Material).Updates(newMaterial).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var MaterialDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.Material{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}
