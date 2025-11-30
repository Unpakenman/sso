package bootstrap

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	authGrpc "sso/internal/app/grpcserver/sso"
)

type GrpcApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int32
}

func NewGrpcServer(
	log *slog.Logger,
	port int32) *GrpcApp {
	gRPCServer := grpc.NewServer()
	authGrpc.Register(gRPCServer)
	return &GrpcApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *GrpcApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *GrpcApp) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	a.log.Info("grpc server started", slog.Int("port", int(a.port)))

	if err := a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve error: %w", err)
	}

	return nil
}

func (a *GrpcApp) Stop() {
	a.log.Info("grpc server stopping")
	a.gRPCServer.GracefulStop()
}
