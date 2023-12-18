package database

import (
	"lpt/pkg/models"

	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

type Database interface {
	AddCourse(models.Course) (int, error)
	AddTopic(models.Topic) (int, error)
	AddExercise(models.Exercise) (int, error)
	GetCourses()
}
