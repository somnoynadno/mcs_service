package API

import (
	"github.com/gorilla/mux"
	"mcs_service/controllers/auth"
	"net/http"
)

func InitAuthRouter(router *mux.Router) {
	router.HandleFunc("/login", auth.Login).Methods(http.MethodPost, http.MethodOptions)
}

