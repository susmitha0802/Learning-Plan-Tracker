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

	if res.RowsAffected == 0 {
		return 0, res.Error
	}

	return int32(topic.ID), nil
}

func (db DBClient) AddExercise(exercise models.Exercise) (int32, error) {
	res := db.DB.Create(&exercise)

	if res.RowsAffected == 0 {
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
