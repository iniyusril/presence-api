package config

import (
	"fmt"
	"os"

	"github.com/iniyusril/presence-api/entity"
	"github.com/iniyusril/presence-api/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	var (
		host        = os.Getenv("DATABASE_HOST")
		user        = os.Getenv("DATABASE_USER")
		password    = os.Getenv("DATABASE_PASSWORD")
		dbname      = os.Getenv("DATABASE_DBNAME")
		port        = os.Getenv("DATABASE_PORT")
		sslMode     = os.Getenv("DATABASE_SSLMODE")
		timezone    = os.Getenv("DATABASE_TIMEZONE")
		automigrate = os.Getenv("DATABASE_AUTOMIGRATE")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslMode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.PanicIfError(err)

	if automigrate == "true" {
		db.AutoMigrate(&entity.User{})
	}

	return db
}
