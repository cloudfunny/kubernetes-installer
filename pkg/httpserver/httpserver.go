package httpserver

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr        string
	routerGroup *gin.RouterGroup
	Engine      *gin.Engine
}

type method string

func InitServer(ip string, port uint32) *Server {
	addr := fmt.Sprintf("%s:%d", ip, port)
	router := gin.Default()
	v1aph1Group := router.Group("/v1alpha1")
	return &Server{
		Addr:        addr,
		routerGroup: v1aph1Group,
		Engine:      router,
	}
}

func (s *Server) addFunc(path string, method string, hander gin.HandlerFunc) {
	switch method {
	case "GET":
		s.routerGroup.GET(path, hander)
	case "POST":
		s.routerGroup.POST(path, hander)
	case "DELETE":
		s.routerGroup.DELETE(path, hander)
	default:
		log.Fatalf("method %v are not supported", method)
	}
}

func (s *Server) Run() error {
	err := s.Engine.Run(s.Addr)
	return err
}
