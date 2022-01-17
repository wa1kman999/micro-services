package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/wa1kman999/helloWorld/handler"
	pb "github.com/wa1kman999/helloWorld/proto"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName    = "helloWorld"
	serviceVersion = "latest"
)

func main() {
	// Create service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.Registry(consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"))),
	)

	service.Init()

	// Register handler
	pb.RegisterHelloWorldHandler(service.Server(), new(handler.UserLogin))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
