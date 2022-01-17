package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/wa1kman999/howAreYou/handler"
	pb "github.com/wa1kman999/howAreYou/proto"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "howareyou"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"))),
	)
	srv.Init()

	// Register handler
	pb.RegisterHowAreYouHandler(srv.Server(), new(handler.HowAreYou))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
