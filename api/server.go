package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minhtam3010/simplebank/db/handler"
	db "github.com/minhtam3010/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	handler := handler.NewAccountHandler(store)

	router.GET("/account/:id", handler.GetAccount)
	router.GET("/account", handler.ListAccount)

	router.POST("/accounts", handler.CreateAccount)

	router.PUT("/account", handler.UpdateAccount)

	router.DELETE("/account", handler.DeleteAccount)
	server.router = router
	return server
}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
