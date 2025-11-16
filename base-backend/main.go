package main

import (
	"fmt"
	"github.com/yadav-shubh/base-backend/config"
	"github.com/yadav-shubh/base-backend/generated/grpc/modules/health/grpc"
	healthGrpc "github.com/yadav-shubh/base-backend/modules/health/server"
	"github.com/yadav-shubh/base-backend/utils"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	utils.GetDB()

	connStr := config.Get().Server.Host
	utils.Log.Info("Binding gRPC server", zap.String("address", connStr))

	listener, err := createReusePortListener("tcp", connStr)
	if err != nil {
		utils.Log.Error("Unable to resolve tcp", zap.String("address", connStr))
	}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	health.RegisterHealthServer(grpcServer, healthGrpc.NewHealthServer())

	err = grpcServer.Serve(listener)
	if err != nil {
		utils.Log.Error("Unable to serve", zap.String("address", connStr))
	}
	utils.Log.Info("Server started", zap.String("address", connStr))

}

func createReusePortListener(network, address string) (net.Listener, error) {
	// Resolve the address
	resolvedAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve address: %w", err)
	}

	// Create a TCP listener
	listener, err := net.ListenTCP(network, resolvedAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create TCP listener: %w", err)
	}

	// Set SO_REUSEADDR and SO_REUSEPORT options
	file, err := listener.File()
	if err != nil {
		return nil, fmt.Errorf("failed to get listener file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			utils.Log.Error("failed to close listener file", zap.Error(err))
		}
	}(file)

	if err := unix.SetsockoptInt(int(file.Fd()), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
		return nil, fmt.Errorf("failed to set SO_REUSEADDR: %w", err)
	}
	if err := unix.SetsockoptInt(int(file.Fd()), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1); err != nil {
		return nil, fmt.Errorf("failed to set SO_REUSEPORT: %w", err)
	}

	return listener, nil
}
