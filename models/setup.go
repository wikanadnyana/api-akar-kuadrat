package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open(mysql.Open("projec27_akar:akarakarakar12@tcp(103.123.62.41:3306)/projec27_akarkuadrat"))
	// database, err := gorm.Open(mysql.Open("root:appengine-laravel@unix('/cloudsql/sacred-highway-400710:asia-southeast2:appengine-laravel')/appengine-laravel?parseTime=true"))
	database, err := gorm.Open(mysql.Open("root:appengine-laravel@tcp(localhost:3306)/appengine-laravel?parseTime=true"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Calculate{})
	DB = database

}
