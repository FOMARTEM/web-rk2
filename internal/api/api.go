package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}

	api.server = echo.New()
	api.server.POST("/answer", api.CreateAnswer)
	api.server.GET("/getQuestions", api.GetQuestions)
	api.server.GET("/getScore", api.GetScore)
	api.server.GET("/newGame", api.newGame)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
