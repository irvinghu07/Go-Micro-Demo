package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	fibo "fibo/proto"
)

type Fibo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Fibo) Cal(ctx context.Context, req *fibo.FiboRequest, rsp *fibo.FiboResponse) error {
	log.Infof("Received Fibo.Cal request: %s\n", req.Rank)
	rsp.Result = fibonacci(int(req.Rank))
	log.Infof("Calculated Result: %v\n", rsp.Result)
	return nil
}

func fibonacci(rank int) int64 {
	if rank < 2 {
		return int64(rank)
	}
	l := make([]int, rank+1)
	l[0] = 0
	l[1] = 1
	for i := 2; i < rank+1; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	log.Info(l)
	return int64(l[rank])
}
