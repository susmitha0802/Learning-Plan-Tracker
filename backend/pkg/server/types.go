package server

import "lpt/pkg/models"

type Database interface {
	AddCourse(models.Course) (int32, error)
	GetCourseNameById(int32) (string, error)
	AddTopic(models.Topic) (int32, error)
	AddExercise(models.Exercise) (int32, error)
	ListCourses()
	ListExerciseIds(int32) ([]int32, error)
	ListSubmittedExercises(int32) ([]models.SubmittedExercises, error)
	AddUser(models.User) (int32, error)
	GetUserDetails(int32) (models.User, error)
	// GetUserEmail(int32) (string, error)
	ListUsersByRole(int32) ([]string, error)
	CreateAssignment(models.CoursesAssigned) (int32, error)
	ListCurrentAssignments() ([]models.CoursesAssigned, error)
	GetUserIdByEmail(string) (int32, error)
	ListAssignedCourses(string) ([]int32, error)
	GetAssignedCourseDetailsByCourseId(int32) (models.Course, error)
	GetAssignedMentorDetails(int32, string) (*models.User, error)
	SubmitExercise(models.SubmittedExercises) (int32, error)
	DeleteExercise(int32, int32) (string, error)
	GetSubmittedExercise(int32, int32) (string, string, error)
	GetProgress(int32, int32) (int32, error)
	ListAssignedMenteesAndCourses(int32) ([]string, []int32, []int32, error)
	ListSubmittedExercisesByMentee(int32, int32) ([]models.ListSubmittedExercisesByMentee, error)
}
