package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB

var err error

func Connect() {
	var dsn = "root:@tcp(localhost:3306)/customer?parseTime=True"

	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
