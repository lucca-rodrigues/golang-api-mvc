package migrations

import (
	"github.com/luccarodrigues/golang-api-mvc/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}