package main

import (
	"github.com/cloudfunny/kubernetes-installer/pkg/httpserver"
	_ "github.com/cloudfunny/kubernetes-installer/pkg/model"
)

func main() {
	// new server
	server := httpserver.InitServer("127.0.0.1", 8080)
	server.RegisterHandler()
	server.Run()
}
