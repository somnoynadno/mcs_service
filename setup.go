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
	createDefaultSectionTypes()
	createDefaultLessonTypes()
}

func createDefaultSectionTypes() {
	for _, st := range entities.DefaultSectionTypes {
		sectionType := entities.SectionType{}
		err := db.GetDB().Where("name = ?", st).First(&sectionType).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				sectionType.Name = st

				err := db.GetDB().Create(&sectionType).Error
				if err != nil {
					panic(err)
				}

				log.Info("Section type '" + st + "' created successfully")
			} else {
				panic(err)
			}
		} else {
			log.Info("Section type '" + st + "' already exists")
		}
	}
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

func createDefaultLessonTypes() {
	for _, lt := range entities.DefaultLessonTypes {
		lessonType := entities.LessonType{}
		err := db.GetDB().Where("name = ?", lt).First(&lessonType).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				lessonType.Name = lt

				err := db.GetDB().Create(&lessonType).Error
				if err != nil {
					panic(err)
				}

				log.Info("Lesson type '" + lt + "' created successfully")
			} else {
				panic(err)
			}
		} else {
			log.Info("Lesson type '" + lt + "' already exists")
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

