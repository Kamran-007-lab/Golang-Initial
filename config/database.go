
package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	"server/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=localhost user=postgres password=Kammu@12345 dbname=mydb port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    } else {
        log.Println("Database connected successfully!")
    }
	DB.AutoMigrate(&models.User{}, &models.Team{})
}
