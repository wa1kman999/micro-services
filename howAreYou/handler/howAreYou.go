package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "github.com/wa1kman999/howAreYou/proto"
)

type HowAreYou struct{}

func (e *HowAreYou) GetUserInfo(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Infof("Received HowAreYou.Call request: %v", req)
	rsp.Name = "小蒋"
	rsp.Age = 18

	return nil
}
