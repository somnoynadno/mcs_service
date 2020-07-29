package API

import (
	"github.com/gorilla/mux"
	"mcs_service/controllers/stats"
	"net/http"
)

func InitStatsRouter(router *mux.Router) {
	router.HandleFunc("/count", stats.GetEntityCount).Methods(http.MethodGet, http.MethodOptions)
}
