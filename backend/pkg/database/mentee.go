package database

import (
	"errors"
	"log"
	"lpt/pkg/models"
)

func (db DBClient) ListAssignedCourses(userEmail string) ([]int32, error) {
	c := []models.CoursesAssigned{}

	menteeId, err := db.GetUserIdByEmail(userEmail)

	if err != nil {
		return nil, errors.New("User Id not found")
	}

	res := db.DB.
		Preload("Mentee").
		Preload("Course").
		Where("mentee_id = ?", menteeId).
		Find(&c)

	coursesId := []int32{}

	for _, v := range c {
		coursesId = append(coursesId, v.CourseId)
	}

	return coursesId, res.Error
}

func (db DBClient) GetAssignedCourseDetailsByCourseId(courseId int32) (models.Course, error) {
	c := models.Course{}
	res := db.DB.
		Preload("Topic").
		Preload("Topic.Exercise").
		Where("id = ?", courseId).
		Find(&c)

	return c, res.Error
}

func (db DBClient) GetAssignedCourseAndMentorDetails(courseId int32, menteeEmail string) (string, error) {
	c := models.CoursesAssigned{}

	menteeId, err := db.GetUserIdByEmail(menteeEmail)

	if err != nil {
		return "", errors.New("User Id not found")
	}

	res := db.DB.
		Preload("Mentor").
		Where("mentee_id = ? AND course_id = ?", menteeId, courseId).
		Find(&c)

	if res.RowsAffected == 0 || res.Error != nil {
		return "", res.Error
	}
	return c.Mentor.Email, res.Error
}

func (db DBClient) SubmitExercise(submit_exercise models.SubmittedExercises) (int32, error) {
	res := db.DB.Create(&submit_exercise)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(submit_exercise.ID), nil
}

func (db DBClient) DeleteExercise(mentee_id int32, exercise_id int32) (string, error) {
	res := db.DB.
		Where("mentee_id = ? AND exercise_id = ?", mentee_id, exercise_id).
		Delete(&models.SubmittedExercises{})

	if res.RowsAffected == 0 {
		if res.Error == nil {
			return "Already deleted / No record found to delete", nil
		}
	}

	if res.Error != nil {
		log.Println("Error", res.Error)
		return "", res.Error
	}

	return "Deleted successfully", nil
}
