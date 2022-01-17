package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/wa1kman999/helloWorld-client/router"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"

	"github.com/gin-gonic/gin"
	log "go-micro.dev/v4/logger"
)

var (
	service = "helloWorld_client"
	version = "latest"
)

func main() {
	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	srv := web.NewService(
		web.Name(service),
		// 需要注册中心
		web.Registry(consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"))),
		web.Version(version),
		web.Address(":9090"),
		web.Handler(r),
	)

	// 初始化
	srv.Init()

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
