package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {
	dbURL := "postgres://postgres:postgres@localhost:5432/pkmdb"
	d, err := gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
