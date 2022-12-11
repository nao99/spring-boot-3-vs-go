package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/config"
)

type Server struct {
	httpServer *fiber.App
}

func NewServer() *Server {
	app := fiber.New()
	return &Server{httpServer: app}
}

func (s *Server) Run(cfg *config.ServerConfig) error {
	portToListen := fmt.Sprintf(":%d", cfg.Port)
	return s.httpServer.Listen(portToListen)
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown()
}

func (s *Server) Get(path string, handler fiber.Handler) {
	s.httpServer.Get(path, handler)
}

func (s *Server) Post(path string, handler fiber.Handler) {
	s.httpServer.Post(path, handler)
}
