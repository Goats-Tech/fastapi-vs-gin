package mygorm

import (
	"fmt"
	"gin/random"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Name  string
	Value float64
}

func Connect(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

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

func Read(c *gin.Context, db *gorm.DB) {
	var product []Product
	result := db.Find(&product)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the product as JSON
	c.JSON(http.StatusOK, product)
}

func Write(c *gin.Context, db *gorm.DB) {
	name := random.GenerateRandomString(10)
	value := random.RandFloat()

	product := Product{Name: name, Value: value}
	db.Create(&product)

	c.JSON(http.StatusOK, product)
}
