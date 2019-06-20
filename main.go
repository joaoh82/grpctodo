package main

import (
	"github.com/joaoh82/shelltodo/pkg/grpc"
	"github.com/joaoh82/shelltodo/pkg/httprouter"
)

func main() {
	// Running gRPC Server on a new goroutine, so it does not block the application
	go grpc.RunServer()

	r := httprouter.SetupRouter()
	r.Run(":3000")
}
