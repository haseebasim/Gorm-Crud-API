package db

import (
	"gorm-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func DbConnection(connectionString string) {
	var err error
	Db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")

}

func Migrate() {
	Db.AutoMigrate(&models.Users{})
	log.Println("Database Migration Completed...")
}
