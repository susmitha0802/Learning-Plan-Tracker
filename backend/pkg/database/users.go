package database

import (
	"log"
	"lpt/pkg/models"
	"lpt/pkg/proto"
)

func (db DBClient) AddUser(user models.User) (int, error) {
	res := db.DB.Create(&user)

	if res.RowsAffected == 0 {
		return 0, res.Error
	}

	return int(user.ID), nil
}

func (db DBClient) GetUsersByRole(role proto.Role) ([]string, error) {
	user_names := []string{}
	res := db.DB.
		Table("users").
		Select("users.name").
		Where("users.role = ?", role).
		Find(&user_names)

	for _, n := range user_names {
		log.Println(n)
	}

	return (user_names), res.Error
}
