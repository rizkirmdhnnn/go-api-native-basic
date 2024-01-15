package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-api-native-basic/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME, "%2F")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Category{}, &models.Member{}, &models.Admin{}, &models.Transactions{}); err != nil {
		panic("failed to migrate database")
	}

	DB = db
	log.Println("Database connected")
}
