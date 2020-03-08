package user

import "github.com/jinzhu/gorm"


func createTables(db *gorm.DB) {
	db.HasTable(user{})
	db.CreateTable(&user{})
}

