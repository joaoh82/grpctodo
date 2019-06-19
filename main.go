package main

import (
	"github.com/joaoh82/shelltodo/pkg/grpc"
	"github.com/joaoh82/shelltodo/pkg/httprouter"
)

func main() {
	go grpc.RunServer()

	r := httprouter.SetupRouter()
	r.Run(":3000")
}
