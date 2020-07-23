package main

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mcs_service/db"
	"mcs_service/models/entities"
	u "mcs_service/utils"
	"os"
)

func createDefaultRolesAndEntities() {
	adminUsername := "admin"
	adminPassword := os.Getenv("default_admin_password")
	if adminPassword == "" {
		panic("no default_admin_password specified")
	}

	// creating admin role
	role := entities.UserRole{}
	err := db.GetDB().Table("user_roles").Select("name = ?", "admin").First(&role).Error

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
		log.Info("Admin role already exist")
	}

	// creating default admin profile
	admin := entities.User{}
	err = db.GetDB().Table("users").Select("username = ?", adminUsername).First(&admin).Error

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
		log.Info("Default admin profile already exist")
	}
}

