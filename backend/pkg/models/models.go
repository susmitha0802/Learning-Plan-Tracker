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
	Time    int32
	Topic   []Topic `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Topic struct {
	gorm.Model
	Name     string
	Resource string
	CourseId int32
	Exercise []Exercise `gorm:"foreignKey:TopicId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Exercise struct {
	gorm.Model
	Question string
	TopicId  int32
}

type User struct {
	gorm.Model
	Name  string
	Email string     `gorm:"unique"`
	Role  proto.Role `gorm:"type:int"`
}

type Tabler interface {
	TableName() string
}

func (CoursesAssigned) TableName() string {
	return "courses_assigned"
}

type CoursesAssigned struct {
	gorm.Model
	MentorId int32
	MenteeId int32
	CourseId int32
	Mentor   User   `gorm:"foreignKey:MentorId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Mentee   User   `gorm:"foreignKey:MenteeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Course   Course `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type SubmittedExercises struct {
	gorm.Model
	MenteeId   int32
	ExerciseId int32
	FileName   string
	File       string
	Mentee     User     `gorm:"foreignKey:MenteeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
