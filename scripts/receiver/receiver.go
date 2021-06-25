package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/profile/middleware"
	"github.com/profile/notifier/generated"
)

type server struct {
	generated.UnimplementedNotifierServer
}

func (s *server) Push(ctx context.Context, in *generated.NotificationRequest) (*generated.NotificationReply, error) {
	_, logger := middleware.LoggerFromContext(ctx)

	logger.Infof("Received: operation %q id %q", in.GetOperation(), in.GetId())
	return &generated.NotificationReply{Message: "ok"}, nil
}

// For notification manual testing
func main() {
	_, logger := middleware.LoggerFromContext(context.Background())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.WithError(err).Fatalln("failed to listen")
	}
	s := grpc.NewServer()
	generated.RegisterNotifierServer(s, &server{})
	logger.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.WithError(err).Fatalln("failed to serve")
	}
}
