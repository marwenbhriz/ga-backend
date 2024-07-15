package models

import (
	"fmt"

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

	//database_host := os.Getenv("DATABASE_HOST")
	//database_port := os.Getenv("DATABASE_PORT")
	//database_user := os.Getenv("DATABASE_USER")
	//database_password := os.Getenv("DATABASE_PASSWORD")

	database, err := gorm.Open(mysql.Open("gauser:pObWTxMQC2T6mSI7@tcp(34.146.152.236:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"))

	//database, err := gorm.Open(mysql.Open(database_user + ":" + database_password + "@tcp(" + database_host + ":" + database_port + ")/books?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Book{})

	DB = database
}
