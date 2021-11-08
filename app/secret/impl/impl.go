package impl

import (
	"database/sql"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/host"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/conf"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db   *sql.DB
	log  logger.Logger
	host host.ServiceServer
	secret.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	s.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (s *service) Name() string {
	return secret.AppName
}

func (s *service) Registry(server *grpc.Server) {
	secret.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
