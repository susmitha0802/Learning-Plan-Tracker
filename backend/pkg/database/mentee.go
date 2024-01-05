package database

import (
	"errors"
	"log"
	"lpt/pkg/models"
	"slices"
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

func (db DBClient) GetAssignedMentorDetails(courseId int32, menteeEmail string) (*models.User, error) {
	c := models.CoursesAssigned{}

	menteeId, err := db.GetUserIdByEmail(menteeEmail)

	if err != nil {
		return nil, errors.New("User Id not found")
	}

	res := db.DB.
		Preload("Mentor").
		Where("mentee_id = ? AND course_id = ?", menteeId, courseId).
		Find(&c)

	if res.RowsAffected == 0 || res.Error != nil {
		return nil, res.Error
	}
	return &c.Mentor, res.Error
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

	if res.Error != nil {
		log.Println("Error", res.Error)
		return "", res.Error
	}

	log.Println(res.RowsAffected)
	if res.RowsAffected == 0 {
		if res.Error == nil {
			return "Already deleted / No record found to delete", nil
		}
	}

	return "Deleted successfully", nil
}

func (db DBClient) GetSubmittedExercise(mentee_id int32, exercise_id int32) (string, string, error) {
	submitted_exercises := models.SubmittedExercises{}
	res := db.DB.
		Where("mentee_id = ? AND exercise_id = ?", mentee_id, exercise_id).
		Find(&submitted_exercises)

	if res.Error != nil {
		log.Println("Error", res.Error)
		return "", "", res.Error
	}

	return submitted_exercises.FileName, submitted_exercises.File, nil
}

func (db DBClient) ListSubmittedExercises(mentee_id int32) ([]models.SubmittedExercises, error) {
	submittedExercises := []models.SubmittedExercises{}
	res := db.DB.
		Preload("Exercise").
		Where("mentee_id = ?", mentee_id).
		Find(&submittedExercises)

	log.Println(submittedExercises, res.RowsAffected)

	if res.Error != nil {
		log.Println("Error", res.Error)
		return nil, res.Error
	}

	if len(submittedExercises) == 0 {
		return nil, nil
	}

	return submittedExercises, nil
}

func (db DBClient) GetProgress(mentee_id int32, course_id int32) (int32, error) {
	exercise_Ids, err := db.ListExerciseIds(course_id)
	if err != nil {
		return 0, err
	}

	total := len(exercise_Ids)
	log.Println("total", total)

	submittedExercises, err := db.ListSubmittedExercises(mentee_id)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, submittedExercise := range submittedExercises {
		if slices.Contains(exercise_Ids, submittedExercise.ExerciseId) {
			count = count + 1

		}
	}

	return int32(count * 100 / total), nil
}
