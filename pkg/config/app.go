package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:vansh@tcp(127.0.0.1:3306)/simpleinterest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
