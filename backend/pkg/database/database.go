package database

import (
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}
