package modules

import (
	"context"
	"github.com/yadav-shubh/base-middleware/generated/grpc/proto"
	"github.com/yadav-shubh/base-middleware/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

var grpcConn *grpc.ClientConn
var oneConn = &sync.Once{}

func getGrpcConn(ctx context.Context) *grpc.ClientConn {
	oneConn.Do(func() {
		addr := "0.0.0.0:8000"
		dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		client, err := grpc.NewClient(addr, dialOptions...)
		if err != nil {
			utils.Log.Error("failed to connect: ", zap.Error(err))
		}
		grpcConn = client
	})
	return grpcConn
}

func getHealthCheckClient(ctx context.Context) health.HealthClient {
	healthClient := getGrpcConn(ctx)
	if healthClient == nil {
		utils.Log.Error("failed to connect for health check")
	}
	return health.NewHealthClient(healthClient)
}
