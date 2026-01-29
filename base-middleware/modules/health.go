package modules

import (
	"context"
	"github.com/yadav-shubh/base-middleware/graph/model"
	"github.com/yadav-shubh/base-middleware/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"time"
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

func CheckHealthSubs(ctx context.Context) <-chan *model.HealthResponse {
	utils.Log.Info("starting health subscription")

	ch := make(chan *model.HealthResponse, 1)

	// run producer goroutine
	go func() {
		defer close(ch)

		ticker := time.NewTicker(2 * time.Second) // send update every 2 seconds
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				utils.Log.Info("subscription closed by client")
				return

			case <-ticker.C:
				// generate new random number every tick
				number := rand.Int()
				parity := "DOWN"
				if number%2 == 0 {
					parity = "UP"
				}

				utils.Log.Info("generated health", zap.Int("number", number))

				// non-blocking send: avoid goroutine freeze
				select {
				case ch <- &model.HealthResponse{Status: parity}:
				default:
					// if client is slow, drop event
				}
			}
		}
	}()

	return ch
}
