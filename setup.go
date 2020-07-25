package main

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"os"
)

func CreateDefaults() {
	createDefaultRolesAndUsers()
	createDefaultTaskTypes()
}

func createDefaultTaskTypes() {
	for _, tt := range entities.DefaultTaskTypes {
		taskType := entities.TaskType{}
		err := db.GetDB().Where("name = ?", tt).First(&taskType).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				taskType.Name = tt

				err := db.GetDB().Create(&taskType).Error
				if err != nil {
					panic(err)
				}

				log.Info("Task type '" + tt + "' created successfully")
			} else {
				panic(err)
			}
		} else {
			log.Info("Task type '" + tt + "' already exists")
		}
	}
}

func createDefaultRolesAndUsers() {
	adminUsername := "admin"
	adminPassword := os.Getenv("default_admin_password")
	if adminPassword == "" {
		panic("no default_admin_password specified")
	}

	// creating admin role
	role := entities.UserRole{}
	err := db.GetDB().Where("name = ?", "admin").First(&role).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			role.Name = "admin"
			role.Description = "Full access to service management"

			err := db.GetDB().Create(&role).Error
			if err != nil {
				panic(err)
			}

			log.Info("Admin role created successfully")
		} else {
			panic(err)
		}
	} else {
		log.Info("Admin role already exists")
	}

	// creating default admin profile
	admin := entities.User{}
	err = db.GetDB().Where("username = ?", adminUsername).First(&admin).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			admin.Username = adminUsername
			admin.Password, _ = u.HashPassword(adminPassword)
			admin.UserRoleID = role.ID

			err := db.GetDB().Create(&admin).Error
			if err != nil {
				panic(err)
			}

			log.Info("Default admin profile created successfully")
		} else {
			panic(err)
		}
	} else {
		log.Info("Default admin profile already exists")
	}
}

