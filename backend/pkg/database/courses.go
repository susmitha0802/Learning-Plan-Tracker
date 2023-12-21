package database

import (
	"log"
	"lpt/pkg/models"
)

func (db DBClient) AddCourse(course models.Course) (int32, error) {
	res := db.DB.Create(&course)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(course.ID), nil
}

func (db DBClient) AddTopic(topic models.Topic) (int32, error) {
	res := db.DB.Create(&topic)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(topic.ID), nil
}

func (db DBClient) AddExercise(exercise models.Exercise) (int32, error) {
	res := db.DB.Create(&exercise)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(exercise.ID), nil
}

func (db DBClient) ListCourses() {
	courses := []models.Course{}
	db.DB.
		Preload("Topic").
		Preload("Topic.Exercise").
		Find(&courses)
	for _, c := range courses {
		log.Println(c)
		for _, t := range c.Topic {
			log.Println(t)
			log.Println("exerices", t.Exercise)
			for _, e := range t.Exercise {
				log.Println(e)
			}
		}
		log.Println()
	}
}

func (db DBClient) ListExerciseIds(course_id int32) ([]int32, error) {
	course := models.Course{}
	res := db.DB.
		Preload("Topic").
		Preload("Topic.Exercise").
		Where("id = ?", course_id).
		Find(&course)

	if res.RowsAffected == 0 || res.Error != nil {
		return nil, res.Error
	}

	exercise_Ids := []int32{}

	for _, topic := range course.Topic {
		for _, exercise := range topic.Exercise {
			exercise_Ids = append(exercise_Ids, int32(exercise.ID))
		}
	}

	return exercise_Ids, nil
}

func (db DBClient) ListSubmittedExercises(mentee_id int32) ([]models.SubmittedExercises, error) {
	submittedExercises := []models.SubmittedExercises{}
	res := db.DB.
		Preload("Exercise").
		Where("mentee_id = ?", mentee_id).
		Find(&submittedExercises)

	if res.RowsAffected == 0 {
		if res.Error == nil {
			return nil, nil
		}
	}

	if res.Error != nil {
		log.Println("Error", res.Error)
		return nil, res.Error
	}

	return submittedExercises, nil
}
