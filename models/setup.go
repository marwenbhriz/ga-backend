package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("no env file exist %v", err)
	}

	database_host := os.Getenv("DATABASE_HOST")
	//database_port := os.Getenv("DATABASE_PORT")
	database_user := os.Getenv("DATABASE_USER")
	database_password := os.Getenv("DATABASE_PASSWORD")

	database, err := gorm.Open(mysql.Open(database_user + ":" + database_password + "@tcp(" + database_host + ")/users-db?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}
