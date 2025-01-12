package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("msyql", "nick:test123/simplerest?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("%s", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
