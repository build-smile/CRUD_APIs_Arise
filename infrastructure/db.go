package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func ConnectDb() {
	// Replace 'user:password@tcp(host:port)/database' with your MySQL connection details
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = CFG.DB.MyDb.Host
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mydb?charset=utf8mb4&parseTime=True&loc=Local",
		CFG.DB.MyDb.User, CFG.DB.MyDb.Password, host, CFG.DB.MyDb.Port)
	// Open a connection to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s err: %v", dsn, err)
	}
	log.Print("db connected")
	DB = db
}
