package database

import (
	"errors"
	"log"
	"lpt/pkg/models"
)

func (db DBClient) AddUser(user models.User) (int32, error) {
	res := db.DB.Create(&user)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(user.ID), nil
}

func (db DBClient) GetUserDetails(userId int32) (models.User, error) {
	user := models.User{}
	res := db.DB.
		Where("id = ?", userId).
		Find(&user)

	if res.RowsAffected == 0 {
		return user, errors.New("There is no user")
	}

	return user, res.Error
}

func (db DBClient) GetUserEmail(userId int32) (string, error) {
	var user_email string
	res := db.DB.
		Table("users").
		Select("email").
		Where("id = ?", userId).
		Find(&user_email)

	if res.RowsAffected == 0 {
		return "", errors.New("There is no user")
	}

	return user_email, res.Error
}

func (db DBClient) ListUsersByRole(roleId int32) ([]string, error) {
	user_names := []string{}
	res := db.DB.
		Table("users").
		Select("users.name").
		Where("users.role = ?", roleId).
		Find(&user_names)

	for _, n := range user_names {
		log.Println(n)
	}

	return (user_names), res.Error
}

func (db DBClient) CreateAssignment(a models.CoursesAssigned) (int32, error) {
	res := db.DB.Create(&a)

	if res.RowsAffected == 0 || res.Error != nil {
		return 0, res.Error
	}

	return int32(a.ID), nil
}

func (db DBClient) ListCurrentAssignments() ([]models.CoursesAssigned, error) {
	c := []models.CoursesAssigned{}
	res := db.DB.
		Preload("Mentor").
		Preload("Mentee").
		Preload("Course").
		Find(&c)

	return c, res.Error
}

func (db DBClient) GetUserIdByEmail(userEmail string) (int32, error) {

	var userId int32

	res := db.DB.Table("users").
		Select("id").
		Where("email = ?", userEmail).
		Find(&userId)

	return userId, res.Error
}
