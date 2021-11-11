package impl

import (
	"database/sql"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/application"
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
	db  *sql.DB
	log logger.Logger
	application.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *service) Name() string {
	return application.AppName
}

func (s *service) Registry(server *grpc.Server) {
	application.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
