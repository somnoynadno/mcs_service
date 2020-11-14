package logical

import (
	"encoding/json"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"net/http"
)

var AddTasksToLesson = func(w http.ResponseWriter, r *http.Request) {
	var TaskLessons []entities.TaskLesson
	err := json.NewDecoder(r.Body).Decode(&TaskLessons)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	tx := db.Begin()

	for i, _ := range TaskLessons {
		err = tx.Create(&TaskLessons[i]).Error

		if err != nil {
			tx.Rollback()
			u.HandleInternalError(w, err)
			return
		}
	}

	tx.Commit()
	res, _ := json.Marshal(&TaskLessons)
	u.RespondJSON(w, res)
}
