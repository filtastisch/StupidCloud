package app

import (
	"StupidCloud/config"
	"StupidCloud/src/internal/node"
	"StupidCloud/src/internal/proxy"
	"StupidCloud/src/internal/service"
	"StupidCloud/src/internal/web"
	"context"
	"log/slog"

	"charm.land/ssh"
	"google.golang.org/grpc"
)

type Controller struct {
	logger *slog.Logger

	config *config.Config

	SSHServer  *ssh.Server
	WebServer  *web.Server
	GRPCServer *grpc.Server

	Services *service.Manager
	Nodes    *node.Manager
	Proxies  *proxy.Manager
}

func (c *Controller) Start(ctx context.Context) interface{} {
	c.logger.Info("Starting controller")

	<-ctx.Done()

	c.logger.Info("Shutting down controller")

	return nil
}

func NewController(logger *slog.Logger, config *config.Config) *Controller {
	return &Controller{
		logger:   logger,
		config:   config,
		Services: service.NewManager(),
		Nodes:    node.NewManager(),
		Proxies:  proxy.NewManager(),
	}
}
