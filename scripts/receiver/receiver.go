package main

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/profile/notifier/generated"
)

type server struct {
	generated.UnimplementedNotifierServer
}

func (s *server) Push(ctx context.Context, in *generated.NotificationRequest) (*generated.NotificationReply, error) {
	logrus.Infof("Received: operation %q id %q", in.GetOperation(), in.GetId())
	return &generated.NotificationReply{Message: "ok"}, nil
}

// For notification manual testing
func main() {
	// port from config.yaml
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.WithError(err).Fatalln("failed to listen")
	}
	s := grpc.NewServer()
	generated.RegisterNotifierServer(s, &server{})
	logrus.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logrus.WithError(err).Fatalln("failed to serve")
	}
}
