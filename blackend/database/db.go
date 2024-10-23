package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// ข้อมูลการเชื่อมต่อฐานข้อมูล
	dsn := "root:@tcp(127.0.0.1:3306)/web_resume?charset=utf8mb4&parseTime=True&loc=Local"

	// เปิดการเชื่อมต่อฐานข้อมูลผ่าน GORM
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to MySQL database successfully.")
}
