package api

import (
	"database/sql"
	"gin/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "host=database port=5432 user=postgres dbname=gin password=postgres sslmode=disable"
)

type Server struct {
	Queries *sqlc.Queries
	Router  *gin.Engine
}

func NewServer() *Server {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("error connecting to database")
	}

	store := sqlc.New(conn)

	server := &Server{
		Queries: store,
		Router:  gin.Default(),
	}

	// Accounts routes
	server.Router.POST("/write", server.write)
	server.Router.GET("/read", server.read)

	return server
}

func (server *Server) Start(url string) error {
	return server.Router.Run(url)
}
