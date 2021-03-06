package routers

import (
	"github.com/gorilla/mux"
	"mcs_service/routers/API"
	"mcs_service/routers/middleware"
	u "mcs_service/utils"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Main API routers
	api := router.PathPrefix("/api").Subrouter()

	// Logical API subrouters
	authRouter := api.PathPrefix("/auth").Subrouter()
	adminRouter := api.PathPrefix("/admin").Subrouter()
	studentRouter := api.PathPrefix("/student").Subrouter()
	statsRouter := api.PathPrefix("/stats").Subrouter()

	API.InitAuthRouter(authRouter)
	API.InitAdminRouter(adminRouter)
	API.InitStudentRouter(studentRouter)
	API.InitStatsRouter(statsRouter)

	// Service command
	api.HandleFunc("/ping", u.HandleOptions).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	// Router middleware usage
	// P.S. Do NOT modify the order
	router.Use(middleware.CORS)              // enable CORS headers
	router.Use(middleware.LogPath)           // log HTTP request URI
	router.Use(middleware.LogBody)           // log HTTP body

	adminRouter.Use(middleware.JwtAuthentication) // attach JWT auth middleware

	return router
}