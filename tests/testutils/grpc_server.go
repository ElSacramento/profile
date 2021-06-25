package testutils

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/profile/middleware"
	"github.com/profile/notifier/generated"
)

type server struct {
	generated.UnimplementedNotifierServer

	grpcServer *grpc.Server
}

func (s *server) Push(ctx context.Context, in *generated.NotificationRequest) (*generated.NotificationReply, error) {
	_, logger := middleware.LoggerFromContext(ctx)

	logger.Infof("Received: operation %q id %q", in.GetOperation(), in.GetId())
	return &generated.NotificationReply{Message: "ok"}, nil
}

// For notification manual testing
func startGRPCServer(addr string) *server {
	_, logger := middleware.LoggerFromContext(context.Background())

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.WithError(err).Fatalln("failed to listen")
		return nil
	}
	s := grpc.NewServer()
	srv := server{grpcServer: s}
	generated.RegisterNotifierServer(s, &srv)
	logger.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.WithError(err).Fatalln("failed to serve")
		return nil
	}
	return &srv
}

func (s *server) stopGRPCServer() {
	s.grpcServer.Stop()
}
