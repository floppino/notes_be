package config

import (
	"beelogiq/notes/controllers"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=postgres-service user=postgres password=nicepwd dbname=notes port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateNoteTable(db)
	controllers.InitiateDB(db)
	return db
}
