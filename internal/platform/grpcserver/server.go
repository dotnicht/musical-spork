package grpcserver

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	grpc *grpc.Server
	ln   net.Listener
}

func New(addr string, opts ...grpc.ServerOption) (*Server, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("listen: %w", err)
	}

	kaParams := keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Time:              2 * time.Hour,
		Timeout:           20 * time.Second,
	}

	baseOpts := []grpc.ServerOption{
		grpc.KeepaliveParams(kaParams),
	}
	baseOpts = append(baseOpts, opts...)

	s := grpc.NewServer(baseOpts...)
	return &Server{grpc: s, ln: ln}, nil
}

func (s *Server) GRPC() *grpc.Server { return s.grpc }
func (s *Server) Serve() error       { return s.grpc.Serve(s.ln) }

func (s *Server) Stop(ctx context.Context) error {
	stopped := make(chan struct{})
	go func() {
		s.grpc.GracefulStop()
		close(stopped)
	}()
	select {
	case <-stopped:
		return nil
	case <-ctx.Done():
		s.grpc.Stop()
		return ctx.Err()
	}
}
