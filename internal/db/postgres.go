package db

import (
	"github.com/priyanshu334/tw-bend/internal/module/auth"
	"github.com/priyanshu334/tw-bend/internal/module/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := database.AutoMigrate(
		&auth.User{},
		&user.UserProfile{},
		&user.Address{},
	); err != nil {
		return err
	}
	DB = database
	return nil
}
