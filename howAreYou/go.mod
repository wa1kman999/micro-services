module github.com/wa1kman999/howAreYou

go 1.16

require (
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20220115202627-a612e09a341f
	go-micro.dev/v4 v4.2.1
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
