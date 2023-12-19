package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=susmitha password=gorm dbname=learning_plan",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	// db.Migrator().DropTable(&Course{}, &Topic{}, &Exercise{})
	db.AutoMigrate(&Course{}, &Topic{}, &Exercise{})
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&User{}, &CoursesAssigned{}, &SubmittedExercises{})

	return db
}
