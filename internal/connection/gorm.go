package connection

import (
	"orm/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormOpen() *gorm.DB {
	dsn := config.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
