package main

import (
	"github.com/irvinghu07/go-micro-demo/handler"
	pb "github.com/irvinghu07/go-micro-demo/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("fibo"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterFiboHandler(srv.Server(), new(handler.Fibo))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
