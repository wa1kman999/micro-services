module github.com/wa1kman999/helloWorld-client

go 1.16

require (
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20220111133620-1e2d197c7d5f
	github.com/gin-gonic/gin v1.7.7
	github.com/wa1kman999/helloworld v0.0.0-00010101000000-000000000000
	go-micro.dev/v4 v4.2.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/wa1kman999/helloworld => ../helloworld
