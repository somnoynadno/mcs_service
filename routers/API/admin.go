package API

import (
	"github.com/gorilla/mux"
	"mcs_service/controllers/crud"
	"mcs_service/controllers/logical"
	"net/http"
)

func InitAdminRouter(router *mux.Router) {
	initLessonRouter(router)
	initSectionRouter(router)
	initSubjectRouter(router)
	initTaskRouter(router)
	initTaskTypeRouter(router)
	initSectionTypeRouter(router)
	initMaterialRouter(router)
}

func initLessonRouter(router *mux.Router) {
	router.HandleFunc("/lesson",      crud.LessonCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/lesson/{id}", crud.LessonRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/lesson/{id}", crud.LessonUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/lesson/{id}", crud.LessonDelete).Methods(http.MethodDelete, http.MethodOptions)
}

func initSectionRouter(router *mux.Router) {
	router.HandleFunc("/section",      crud.SectionCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/section/{id}", crud.SectionRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/section/{id}", crud.SectionUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/section/{id}", crud.SectionDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/sections_by_subject/{subject_id}",
		logical.GetSectionsBySubjectID).Methods(http.MethodGet, http.MethodOptions)
}

func initSubjectRouter(router *mux.Router) {
	router.HandleFunc("/subject",      crud.SubjectCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/subject/{id}", crud.SubjectRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/subject/{id}", crud.SubjectUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/subject/{id}", crud.SubjectDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/all_subjects", logical.GetAllSubjects).Methods(http.MethodGet, http.MethodOptions)
}

func initTaskRouter(router *mux.Router) {
	router.HandleFunc("/task",      crud.TaskCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/task/{id}", crud.TaskRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/task/{id}", crud.TaskUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/task/{id}", crud.TaskDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/tasks_by_section/{section_id}",
		logical.GetTasksBySectionID).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/tasks_by_task_type/{task_type_id}",
		logical.GetTasksByTaskTypeID).Methods(http.MethodGet, http.MethodOptions)
}

func initMaterialRouter(router *mux.Router) {
	router.HandleFunc("/material",      crud.MaterialCreate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/material/{id}", crud.MaterialRetrieve).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/material/{id}", crud.MaterialUpdate).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/material/{id}", crud.MaterialDelete).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/materials_by_section/{section_id}",
		logical.GetMaterialsBySectionID).Methods(http.MethodGet, http.MethodOptions)
}

func initLessonTypeRouter(router *mux.Router) {
	router.HandleFunc("/all_lesson_types", logical.GetAllLessonTypes).Methods(http.MethodGet, http.MethodOptions)
}

func initTaskTypeRouter(router *mux.Router) {
	router.HandleFunc("/all_task_types", logical.GetAllTaskTypes).Methods(http.MethodGet, http.MethodOptions)
}

func initSectionTypeRouter(router *mux.Router) {
	router.HandleFunc("/all_section_types", logical.GetAllSectionTypes).Methods(http.MethodGet, http.MethodOptions)
}

