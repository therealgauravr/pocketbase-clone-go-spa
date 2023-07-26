package db

import (
	"log"

	"github.com/therealgauravr/pocketbase-clone-spa/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitORM() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=default password=default dbname=postgres port=8000 sslmode=disable TimeZone=Asia/Kolkata client_encoding=UTF8"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform auto-migration of foundational structures

	db.AutoMigrate(&types.User{})

	return db, nil
}

func GetUsers(db *gorm.DB) []types.User {
	var result []types.User
	err := db.Find(&result).Error
	if err != nil {
		log.Println("No users detected...")
	}
	return result
}

func ShowTables(db *gorm.DB) {
	var result []TableRows
	db.Raw("SELECT * FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND  schemaname != 'information_schema';").Scan(&result)
	log.Println("FOLLOWING TABLES DETECTED")
	log.Printf("%+v", result)
}
