package modules

import (
	"context"
	"github.com/yadav-shubh/base-middleware/graph/model"
	"github.com/yadav-shubh/base-middleware/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

func CheckHealth(ctx context.Context) *model.HealthResponse {
	client := getHealthCheckClient(ctx)
	var r, err = client.HealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		utils.Log.Error("failed to connect for health check")
	}

	if r == nil {
		return &model.HealthResponse{
			Status: "DOWN",
		}
	}

	return &model.HealthResponse{
		Status: r.GetStatus(),
	}
}
