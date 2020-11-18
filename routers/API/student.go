package API

import (
	"github.com/gorilla/mux"
	"mcs_service/controllers/logical"
	"mcs_service/controllers/student"
	"net/http"
)

func InitStudentRouter(router *mux.Router) {
	router.HandleFunc("/all_subjects", logical.GetAllSubjects).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/all_sections", logical.GetAllSections).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/sections_by_subject/{subject_id}",
		student.GetSectionsBySubjectID).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/materials_by_section/{section_id}",
		logical.GetMaterialsBySectionID).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/encrypted_tasks_by_section/{section_id}",
		student.GetTasksBySectionID).Methods(http.MethodGet, http.MethodOptions)
}
