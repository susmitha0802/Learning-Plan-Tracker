package models

import (
	"lpt/pkg/proto"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Caption string
	Logo    string
	Time    int
	Topic   []Topic `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Topic struct {
	gorm.Model
	Name     string
	Resource string
	CourseId int
	Exercise []Exercise `gorm:"foreignKey:TopicId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Exercise struct {
	gorm.Model
	Question string
	TopicId  int
}

type User struct {
	gorm.Model
	Name  string
	Email string     `gorm:"unique"`
	Role  proto.Role `gorm:"type:int"`
}

type CoursesAssigned struct {
	gorm.Model
	MentorId int
	MenteeId int
	CourseId int
}

type SubmittedExercises struct {
	gorm.Model
	MenteeId   int
	ExerciseId int
	FileName   string
	File       string
}
