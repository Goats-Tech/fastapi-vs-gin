package api

import (
	"gin/db/sqlc"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) read(c *gin.Context) {
	rows, err := server.Queries.ReadRows(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rows)

}

func (server *Server) write(c *gin.Context) {
	err := server.Queries.WriteRows(c, sqlc.WriteRowsParams{
		Name:  utils.RandomOwner(),
		Value: int32(utils.RandomMoney()),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, err)
}
