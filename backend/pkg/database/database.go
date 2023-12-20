package database

import (
	"lpt/pkg/models"

	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

type Database interface {
	AddCourse(models.Course) (int32, error)
	AddTopic(models.Topic) (int32, error)
	AddExercise(models.Exercise) (int32, error)
	ListCourses()
	AddUser(models.User) (int32, error)
	ListUsersByRole(int32) ([]string, error)
	CreateAssignment(models.CoursesAssigned) (int32, error)
	ListCurrentAssignments() ([]models.CoursesAssigned, error)
	GetUserIdByEmail(string) (int32, error)
	ListAssignedCourses(string) ([]int32, error)
	GetAssignedCourseDetailsByCourseId(int32) (models.Course, error)
	GetAssignedCourseAndMentorDetails(int32, string) (string, error)
	SubmitExercise(models.SubmittedExercises) (int32, error)
	DeleteExercise(int32, int32) (string, error)
}
