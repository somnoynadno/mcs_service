package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"mcs_service/models/entities"
	"os"
)

var db *gorm.DB

func init() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, username, dbName, password))
	if err != nil {
		panic(err)
	} else {
		db = conn
		log.Info("DB connected on " + dbHost)
	}

	err = migrateSchema()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Schema migrated successfully")
	}
}

func GetDB() *gorm.DB {
	return db
}

func migrateSchema() error {
	err := db.AutoMigrate(
		entities.Lesson{},
		entities.LessonType{},
		entities.Material{},
		entities.Section{},
		entities.SectionType{},
		entities.ServiceLogs{},
		entities.Subject{},
		entities.Task{},
		entities.TaskLesson{},
		entities.TaskType{},
		entities.User{},
		entities.UserRole{},
	).Error

	return err
}
