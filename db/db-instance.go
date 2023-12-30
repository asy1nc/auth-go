package db

import (
	"fmt"

	"github.com/asy1nc/auth-go/config"
	models "github.com/asy1nc/auth-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	if db != nil {
		return db
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

func Migrate(wipe bool) {
	db := Connect()

	if wipe {
		db.Migrator().DropTable(&models.User{}, &models.Password{})
	}

	db.AutoMigrate(&models.User{}, &models.Password{})
}
