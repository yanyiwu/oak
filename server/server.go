package server

import "github.com/yanyiwu/settledb/engine"

type Server struct {
	engine *engine.Engine
}

func NewServer() *Server {
	server := &Server{engine.NewEngine()}
	return server
}

func (s *Server) Start() {
}
