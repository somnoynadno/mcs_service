package API

import (
	"github.com/gorilla/mux"
	"mcs_service/controllers/crud"
	"mcs_service/controllers/logical"
	"net/http"
)

func InitAdminRouter(router *mux.Router) {
	initSectionRouter(router)
	initSubjectRouter(router)
	initTaskRouter(router)
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
}
