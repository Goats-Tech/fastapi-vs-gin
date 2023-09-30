package main

import (
	"gin/db"
	"gin/random"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Name  string
	Value float64
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

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		return
	}

	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/read", func(c *gin.Context) {
		Read(c, dbConn)
	})
	server.POST("/write", func(c *gin.Context) {
		Write(c, dbConn)
	})

	err = server.Run(":5555")
	if err != nil {
		return
	}
}
