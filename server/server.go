package server

import (
	"log"

	"github.com/DanielDDHM/management/config"
	"github.com/DanielDDHM/management/docs"
	"github.com/DanielDDHM/management/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   config.GetConfig().ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	s.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.server.Use(cors.Default())
	router := routes.ConfigRoutes(s.server)

	docs.SwaggerInfo.Title = "management"
	docs.SwaggerInfo.Description = "Version 1 management"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3333"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
