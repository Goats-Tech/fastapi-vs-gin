package sqlc

import (
	"gin/random"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

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
