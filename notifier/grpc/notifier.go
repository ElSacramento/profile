package grpc

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/profile/configuration"
	"github.com/profile/middleware"
	"github.com/profile/notifier"
	"github.com/profile/notifier/generated"
)

type Sender struct {
	clients     []generated.NotifierClient
	connections []*grpc.ClientConn
	logger      *logrus.Entry
}

func New(ctx context.Context, notify configuration.Notify) *Sender {
	_, logger := middleware.LoggerFromContext(ctx)

	clients := make([]generated.NotifierClient, 0)
	connections := make([]*grpc.ClientConn, 0)

	for _, elem := range notify {
		if elem.Addr != "" {
			conn, err := grpc.Dial(elem.Addr, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				logger.WithError(err).Error("could not connect to grpc server")
				continue
			}
			c := generated.NewNotifierClient(conn)
			clients = append(clients, c)
			connections = append(connections, conn)
		}
	}

	return &Sender{clients: clients, connections: connections, logger: logger}
}

func (n *Sender) Push(id uint64, opName notifier.Operation) {
	msg := &generated.NotificationRequest{Id: id, Operation: string(opName)}
	wg := &sync.WaitGroup{}
	for _, client := range n.clients {
		go func(client generated.NotifierClient, logger *logrus.Entry) {
			defer wg.Done()
			wg.Add(1)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := client.Push(ctx, msg)
			if err != nil {
				logger.WithError(err).Error("could not send push notification")
				return
			}
			if r.GetMessage() != "ok" {
				logger.Error("bad response")
				return
			}
		}(client, n.logger)
	}

	wg.Wait()
}

func (n *Sender) Stop() {
	for _, conn := range n.connections {
		if err := conn.Close(); err != nil {
			n.logger.WithError(err).Error("failed to close connection")
		}
	}
}
