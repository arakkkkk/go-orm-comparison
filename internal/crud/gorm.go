package crud

import (
	"fmt"
	"log"

	"orm/internal/connection"
	table "orm/internal/migration"
)

func GoemCreate() {
	db := connection.GormOpen()

	user := table.GormUser{
		Name:     "太郎",
		Age:      20,
		IsActive: true,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// Fingの際は引数がinterface{}になっており、エラー
/*
func (*gorm.DB).Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
Find finds all records matching given conditions conds
[`(gorm.DB).Find` on pkg.go.dev](https://pkg.go.dev/gorm.io/gorm@v1.25.4#DB.Find)
*/
func GooemRead() {
	db := connection.GormOpen()
	users := []table.GormUser{}
	result := db.Find(&users) // 引数がinterface
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}
