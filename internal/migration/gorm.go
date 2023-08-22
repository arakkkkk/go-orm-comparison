package migration

import (
	"gorm.io/gorm"
	"orm/internal/connection"
)

type GormUser struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model
	Name     string
	Age      int
	IsActive bool
}

func GormMigrate() {
	// dbを作成します
	db := connection.GormOpen()
	// dbをmigrateします
	db.AutoMigrate(&GormUser{})
}
