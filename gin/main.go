package main

import (
	"context"
	"gin/mygorm"
	"gin/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	useGorm := true
	connString := "postgresql://postgres:secret@fastapi-vs-gin-database-1:5432/gin?sslmode=disable"

	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if useGorm {
		dbConn, err := mygorm.Connect(connString)
		if err != nil {
			return
		}

		server.GET("/read", func(c *gin.Context) {
			mygorm.Read(c, dbConn)
		})
		server.POST("/write", func(c *gin.Context) {
			mygorm.Write(c, dbConn)
		})
	} else {
		dbConn, err := pgxpool.New(context.Background(), connString)
		if err != nil {
			return
		}

		server.GET("/read", func(c *gin.Context) {
			_, err := sqlc.New(dbConn).Read(c, 100)
			if err != nil {
				return
			}
		})
		server.POST("/write", func(c *gin.Context) {
			_, err := sqlc.New(dbConn).Read(c, 100)
			if err != nil {
				return
			}
		})
	}

	err := server.Run(":5555")
	if err != nil {
		return
	}
}
