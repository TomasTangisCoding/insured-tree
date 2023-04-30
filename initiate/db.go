package initiate

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB

var err error

func ConnectDB() {
	DBconnect, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
