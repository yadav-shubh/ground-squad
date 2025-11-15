package server

import (
	"context"
	pb "github.com/yadav-shubh/base-backend/generated/grpc/modules/health/grpc"
	"github.com/yadav-shubh/base-backend/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthServer struct {
	pb.UnimplementedHealthServer
}

func NewHealthServer() *HealthServer {
	return &HealthServer{
		UnimplementedHealthServer: pb.UnimplementedHealthServer{},
	}
}

func (s *HealthServer) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	utils.Log.Info("Health check called", zap.String("status", "UP"))
	return &pb.HealthCheckResponse{
		Status: "UP",
	}, nil
}
