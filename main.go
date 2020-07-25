package main

import (
	log "github.com/sirupsen/logrus"
	"mcs_service/db"
	"mcs_service/routers"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	Router := routers.InitRouter()

	// check connection
	con := db.GetDB()
	errors := con.GetErrors()
	if errors != nil && len(errors) > 0 {
		panic(errors[0])
	}

	// set things up
	CreateDefaults()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // localhost
	}

	log.Info("Listening on: ", port)
	err := http.ListenAndServe(":" + port, Router)
	if err != nil {
		log.Panic(err)
	}
}
