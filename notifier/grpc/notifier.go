package grpc

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
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

	wg := &sync.WaitGroup{}
	msgChan := make(chan *grpc.ClientConn, len(notify))
	errChan := make(chan error, len(notify))

	logger.Infof("Creating %d clients", len(notify))
	for _, addr := range notify {
		if addr != "" {
			wg.Add(1)

			go func(addr string, wg *sync.WaitGroup) {
				defer wg.Done()

				ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				conn, err := grpc.DialContext(ctxWithTimeout, addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableRetry())
				if err != nil {
					errChan <- errors.Wrap(err, "could not connect to grpc server on "+addr)
					return
				}
				msgChan <- conn
			}(addr, wg)
		}
	}

	wg.Wait()
	close(errChan)
	close(msgChan)

	for val := range errChan {
		logger.Error(val)
	}

	for val := range msgChan {
		c := generated.NewNotifierClient(val)
		clients = append(clients, c)
		connections = append(connections, val)
		logger.Info("GRPC client created")
	}

	return &Sender{clients: clients, connections: connections, logger: logger}
}

func (n *Sender) Push(id uint64, opName notifier.Operation) {
	n.logger.Infof("Send notification to %d listeners", len(n.clients))

	wg := &sync.WaitGroup{}

	errChan := make(chan error, len(n.clients))
	for _, client := range n.clients {
		wg.Add(1)

		go func(client generated.NotifierClient, id uint64, opName string, wg *sync.WaitGroup) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := client.Push(ctx, &generated.NotificationRequest{Id: id, Operation: opName})
			if err != nil {
				errChan <- errors.Wrap(err, "could not send push notification")
				return
			}
			if r.GetMessage() != "ok" {
				errChan <- errors.New("bad response")
				return
			}
		}(client, id, string(opName), wg)
	}

	wg.Wait()
	close(errChan)

	for val := range errChan {
		n.logger.Error(val)
	}
}

func (n *Sender) Stop() {
	for _, conn := range n.connections {
		if err := conn.Close(); err != nil {
			n.logger.WithError(err).Error("failed to close connection")
		}
	}
}
