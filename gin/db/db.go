package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Value float64
}

func Connect() (*gorm.DB, error) {
	dsn := "host=fastapi-vs-gin-database-1 port=5432 user=postgres dbname=gin password=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	Migrate(db)

	return db, nil
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
}
