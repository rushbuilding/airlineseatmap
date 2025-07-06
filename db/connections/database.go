package connections

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func GetConnection() *gorm.DB {
	connect := mysql.Open("seat:123456abc@tcp(localhost:3306)/seatmapdatabase?charset=utf8mb4&parseTime=True&loc=UTC")
	db, err := gorm.Open(connect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		DisableAutomaticPing: true, // optional, improves speed
		PrepareStmt:          false, // optional, enables caching prepared statements
	})
	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)

	return db
}