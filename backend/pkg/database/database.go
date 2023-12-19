package database

import (
	"lpt/pkg/models"
	"lpt/pkg/proto"

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
	AddUser(models.User) (int, error)
	GetUsersByRole(proto.Role) ([]string, error)
}
