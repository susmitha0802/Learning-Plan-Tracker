package database

import (
	"log"
	"lpt/pkg/models"
)

func (db DBClient) AddCourse(course models.Course) (int, error) {
	res := db.DB.Create(&course)

	if res.RowsAffected == 0 {
		return 0, res.Error
	}

	return int(course.ID), nil
}

func (db DBClient) AddTopic(topic models.Topic) (int, error) {
	res := db.DB.Create(&topic)

	if res.RowsAffected == 0 {
		return 0, res.Error
	}

	return int(topic.ID), nil
}

func (db DBClient) AddExercise(exercise models.Exercise) (int, error) {
	res := db.DB.Create(&exercise)

	if res.RowsAffected == 0 {
		return 0, res.Error
	}

	return int(exercise.ID), nil
}

func (db DBClient) GetCourses() {
	courses := []models.Course{}
	db.DB.Model(&courses).Preload("topics").
		Find(&courses)
	for _, c := range courses {
		log.Println(c)
		for _, t := range c.Topic {
			log.Println(t)
		}
		log.Println()
	}
}
