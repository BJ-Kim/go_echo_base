package app

import (
	"github.com/labstack/echo/v4"
	"web_base/app/endpoints"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	return &Server{echo.New()}
}

func (sv *Server) registerEndpoints() {
	ep := endpoints.NewEndpoints(sv.e)
	ep.RegisterEndpoints()
}

func (sv *Server) RunServer() {
	sv.registerEndpoints()
	sv.e.Logger.Fatal(sv.e.Start(":1234"))
}
