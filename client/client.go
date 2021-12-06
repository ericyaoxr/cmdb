package client

import (
	"google.golang.org/grpc"

	"github.com/ericyaoxr/cmdb/app/application"
	"github.com/ericyaoxr/cmdb/app/config"
	"github.com/ericyaoxr/cmdb/app/host"
	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/app/task"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"
)

// NewClient todo
func NewClient(conf *Config) (*Client, error) {
	zap.DevelopmentSetup()

	conn, err := grpc.Dial(
		conf.address,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("CMDB SDK"),
	}, nil
}

// Client 客户端
type Client struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Resource todo
func (c *Client) Resource() resource.ServiceClient {
	return resource.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Secret() secret.ServiceClient {
	return secret.NewServiceClient(c.conn)
}

// Task todos
func (c *Client) Task() task.ServiceClient {
	return task.NewServiceClient(c.conn)
}

// Application todos
func (c *Client) Application() application.ServiceClient {
	return application.NewServiceClient(c.conn)
}

// Config todos
func (c *Client) Config() config.ServiceClient {
	return config.NewServiceClient(c.conn)
}
